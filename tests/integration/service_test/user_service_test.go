package service_test

import (
	"clean-architecture/domain/features/user"
	"clean-architecture/tests/setup"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/fx"
)

func TestUserService(t *testing.T) {
	var s *user.Service

	_, cancel, err := setup.DI(t, fx.Options(fx.Populate(&s)))
	defer cancel()
	if err != nil {
		log.Println(err)
		return
	}

	t.Run("User can be created", func(t *testing.T) {
		_user := user.User{
			Name:  "dipesh",
			Age:   2,
			Email: "dipesh.dulal@wesionary.team",
		}
		err := s.Create(&_user)
		assert.NoError(t, err, "user creation fails")

		gotUser, err := s.GetOneUser(_user.ID)
		assert.NoError(t, err, "user get fails")
		assert.Equal(t, _user.Email, gotUser.Email, "same user returned from db")
	})
}
