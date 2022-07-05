package api_test

import (
	"clean-architecture/infrastructure"
	"clean-architecture/tests/setup"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestUserRoutes(t *testing.T) {
	var router infrastructure.Router

	_, cancel, err := setup.DI(t, fx.Options(fx.Populate(&router)))
	defer cancel()
	if err != nil {
		log.Println(err)
		return
	}

	t.Run("user get error if bad formatted uuid is not provided", func(t *testing.T) {
		testPath := "/api/user/123"
		testMethod := "GET"
		response := httptest.NewRecorder()
		request := httptest.NewRequest(testMethod, testPath, nil)
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusBadRequest, response.Code, "status code match")
	})

	t.Run("user get error if uuid not in database is provided", func(t *testing.T) {
		testPath := "/api/user/32404d15-a878-4392-8ce3-75d4b7e038ce"
		testMethod := "GET"
		response := httptest.NewRecorder()
		request := httptest.NewRequest(testMethod, testPath, nil)
		router.ServeHTTP(response, request)
		assert.Equal(t, http.StatusInternalServerError, response.Code, "status code match")
	})
}
