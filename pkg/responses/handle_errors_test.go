package responses_test

import (
	"clean-architecture/pkg/errorz"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"clean-architecture/pkg/utils"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type MockLogger struct {
	Errors []error
}

func (m *MockLogger) Error(err error) {
	m.Errors = append(m.Errors, err)
}

func TestHandleError(t *testing.T) {
	testCases := []struct {
		name                string
		err                 error
		expectedStatusCode  int
		expectedBody        string
		expectSentryCapture bool
	}{
		{
			name:                "Handle API Error",
			err:                 errorz.ErrBadRequest,
			expectedStatusCode:  http.StatusBadRequest,
			expectedBody:        `{"error":"Bad Request"}`,
			expectSentryCapture: false,
		},
		{
			name:                "Handle API Error With Custom Message",
			err:                 errorz.ErrBadRequest.JoinError("Bad Request"),
			expectedStatusCode:  http.StatusBadRequest,
			expectedBody:        `{"error":"Bad Request"}`,
			expectSentryCapture: false,
		},
		{
			name:                "Handle Already Exists API Error",
			err:                 errorz.ErrAlreadyExists,
			expectedStatusCode:  http.StatusConflict,
			expectedBody:        `{"error":"Already Exists"}`,
			expectSentryCapture: false,
		},
		{
			name:                "Handle Record Not Found Error",
			err:                 gorm.ErrRecordNotFound,
			expectedStatusCode:  http.StatusNotFound,
			expectedBody:        `{"error":"record not found"}`,
			expectSentryCapture: false,
		},
		{
			name:                "Handle Generic Error",
			err:                 errors.New("something went wrong"),
			expectedStatusCode:  http.StatusInternalServerError,
			expectedBody:        `{"error":"An error occurred while processing your request. Please try again later."}`,
			expectSentryCapture: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := &MockSentryService{}
			originalService := utils.CurrentSentryService
			utils.CurrentSentryService = mockService
			defer func() {
				utils.CurrentSentryService = originalService
				mockService.Reset()
			}()

			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Request, _ = http.NewRequest("POST", "/", nil)

			testLogger := framework.CreateTestLogger(t)

			responses.HandleError(testLogger, c, tc.err)
			assert.Equal(t, tc.expectedStatusCode, w.Code)
			assert.JSONEq(t, tc.expectedBody, w.Body.String())

			if tc.expectSentryCapture {
				assert.True(t, mockService.WasCalled(), "Expected Sentry to capture the error")
			} else {
				assert.False(t, mockService.WasCalled(), "Sentry should not capture this error")
			}
		})
	}
}
