package controllers_test

import (
	"clean-architecture/api/controllers"
	"clean-architecture/lib"
	"clean-architecture/tests/setup"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestUserController(t *testing.T) {
	var uc *controllers.UserController

	_, cancel, err := setup.DI(t, fx.Options(fx.Populate(&uc)))
	defer cancel()
	if err != nil {
		log.Println(err)
		return
	}

	t.Run("user get error if bad formatted uuid is not provided", func(t *testing.T) {
		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)

		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: "123",
			},
		}

		uc.GetOneUser(c)

		assert.Equal(t, http.StatusBadRequest, response.Code, "status code match")

	})

	t.Run("user get error if uuid not in database is provided", func(t *testing.T) {
		response := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(response)

		id, _ := uuid.NewRandom()

		c.Params = []gin.Param{
			{
				Key:   "id",
				Value: lib.BinaryUUID(id).String(),
			},
		}

		uc.GetOneUser(c)

		assert.Equal(t, http.StatusInternalServerError, response.Code, "status code match")
	})
}
