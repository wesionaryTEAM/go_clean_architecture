package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"prototype2/domain"
	servMock "prototype2/domain/mocks"
	fbMock "prototype2/api/service/mocks"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddUser(t *testing.T) {
	input := &domain.User{
		Name: "Tester",
		Email: "tester@gmail.com",
		Password: "Zxcvbnm123",
		DOB: "2020-07-25",
	}

	_, err := json.Marshal(&input)
	assert.NoError(t, err)

	mockUserService := new(servMock.UserService)
	mockFbService := new(fbMock.FirebaseService)

	//Setting up the expectations
	mockUserService.On("Validate", &input).Return(nil)
	mockUserService.On("ValidateAge", &input).Return(true)

	input.ID = "uid111"
	mockFbService.On("CreateUser", input.Email, input.Password).Return(input.ID, nil)
	mockUserService.On("Create", &input).Return(&input, nil)

	testController := NewUserController(mockUserService, mockFbService)

	gin.SetMode(gin.TestMode)
	r := gin.Default()
	resp := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(resp)

	c.Request, _ = http.NewRequest("POST", "/posts", nil)
	r.ServeHTTP(resp, c.Request)

	testController.AddUser(c)

	res := resp.Result()
	assert.Equal(t, 200, res.StatusCode)
}

