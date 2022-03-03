package service_test

import (
	"clean-architecture/models"
	"clean-architecture/services"
	"clean-architecture/tests/setup"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestUserService(t *testing.T) {
	var s *services.UserService

	_, cancel, err := setup.SetupDI(t, fx.Options(fx.Populate(&s)))
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}

	t.Run("User can be created", func(t *testing.T) {
		err := s.Create(&models.User{
			Name:  "dipesh",
			Age:   2,
			Email: "dipesh.dulal@wesionary.team",
		})
		assert.NoError(t, err, "user creation fails")

		users, err := s.GetAllUser()
		assert.NoError(t, err, "user get fails")

		fmt.Println(users)
	})
}
