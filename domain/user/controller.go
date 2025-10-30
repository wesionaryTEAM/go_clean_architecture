package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"clean-architecture/pkg/types"

	"github.com/gin-gonic/gin"
)

// UserController data type
type Controller struct {
	service *Service
	logger  framework.Logger
	env     *framework.Env
}

type URLObject struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

// NewUserController creates new user controller
func NewController(
	userService *Service,
	logger framework.Logger,
	env *framework.Env,
) *Controller {
	return &Controller{
		service: userService,
		logger:  logger,
		env:     env,
	}
}

// CreateUser creates the new user
func (u *Controller) CreateUser(c *gin.Context) {
	var user models.User

	if err := c.Bind(&user); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	// check if the user already exists

	if err := u.service.Create(&user); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{"data": "user created"})
}

// GetOneUser gets one user
func (u *Controller) GetUserByID(c *gin.Context) {
	paramID := c.Param("id")

	userID, err := types.ShouldParseUUID(paramID)
	if err != nil {
		responses.HandleValidationError(u.logger, c, ErrInvalidUserID)
		return
	}

	user, err := u.service.GetUserByID(userID)
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}
