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
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

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
			expectedBody:        `{"error":{"code":"BAD_REQUEST","severity":"warn","message":"Bad Request"}}`,
			expectSentryCapture: false,
		},
		{
			name:                "Handle API Error With Custom Message",
			err:                 errorz.ErrBadRequest.WithMessage("custom message"),
			expectedStatusCode:  http.StatusBadRequest,
			expectedBody:        `{"error":{"code":"BAD_REQUEST","severity":"warn","message":"custom message"}}`,
			expectSentryCapture: false,
		},
		{
			name:                "Handle Already Exists API Error",
			err:                 errorz.ErrAlreadyExists,
			expectedStatusCode:  http.StatusConflict,
			expectedBody:        `{"error":{"code":"ALREADY_EXISTS","severity":"warn","message":"Conflict"}}`,
			expectSentryCapture: false,
		},
		{
			name:                "Handle Record Not Found Error",
			err:                 gorm.ErrRecordNotFound,
			expectedStatusCode:  http.StatusNotFound,
			expectedBody:        `{"error":{"code":"RECORD_NOT_FOUND","severity":"warn","message":"Not Found"}}`,
			expectSentryCapture: false,
		},
		{
			name:                "Handle Generic Error",
			err:                 errors.New("something went wrong"),
			expectedStatusCode:  http.StatusInternalServerError,
			expectedBody:        `{"error":{"code":"INTERNAL_ERROR","severity":"error","message":"Internal Server Error"}}`,
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

func TestHandleValidationError_APIErrorPassThrough(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/", nil)

	apiErr := errorz.New("INVALID_USER_ID", http.StatusBadRequest, errorz.SeverityWarn)
	responses.HandleValidationError(framework.CreateTestLogger(t), c, apiErr)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error":{"code":"INVALID_USER_ID","severity":"warn"}}`, w.Body.String())
}

func TestHandleValidationError_OzzoErrors(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/", nil)

	verrs := validation.Errors{"Email": errors.New("cannot be blank")}
	responses.HandleValidationError(framework.CreateTestLogger(t), c, verrs)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t,
		`{"error":{"code":"BAD_REQUEST","severity":"warn","message":"Bad Request","params":{"validation_errors":[{"field":"Email","error_type":"validation","message":"cannot be blank"}]}}}`,
		w.Body.String())
}

func TestHandleValidationError_OzzoErrorsAreSortedByField(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/", nil)

	verrs := validation.Errors{
		"Email":     errors.New("cannot be blank"),
		"FirstName": errors.New("is required"),
	}
	responses.HandleValidationError(framework.CreateTestLogger(t), c, verrs)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	// Output must be deterministically ordered by field regardless of map iteration order.
	assert.JSONEq(t,
		`{"error":{"code":"BAD_REQUEST","severity":"warn","message":"Bad Request","params":{"validation_errors":[`+
			`{"field":"Email","error_type":"validation","message":"cannot be blank"},`+
			`{"field":"FirstName","error_type":"validation","message":"is required"}]}}}`,
		w.Body.String())
}

func TestHandleErrorWithParams(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/", nil)

	responses.HandleErrorWithParams(framework.CreateTestLogger(t), c,
		errorz.ErrAlreadyExists, map[string]any{"email": "a@b.com"})

	assert.Equal(t, http.StatusConflict, w.Code)
	assert.JSONEq(t,
		`{"error":{"code":"ALREADY_EXISTS","severity":"warn","message":"Conflict","params":{"email":"a@b.com"}}}`,
		w.Body.String())
}

func TestHandleErrorWithStatus(t *testing.T) {
	testCases := []struct {
		name             string
		statusCode       int
		expectedSeverity string
		expectSentry     bool
	}{
		{name: "4xx does not capture to Sentry", statusCode: http.StatusBadRequest, expectedSeverity: "warn", expectSentry: false},
		{name: "5xx captures to Sentry", statusCode: http.StatusInternalServerError, expectedSeverity: "error", expectSentry: true},
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

			responses.HandleErrorWithStatus(framework.CreateTestLogger(t), c, tc.statusCode, errors.New("boom"))

			assert.Equal(t, tc.statusCode, w.Code)
			assert.JSONEq(t,
				`{"error":{"code":"CUSTOM_ERROR","severity":"`+tc.expectedSeverity+`"}}`,
				w.Body.String())
			assert.Equal(t, tc.expectSentry, mockService.WasCalled())
		})
	}
}
