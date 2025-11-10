package user

import (
	"clean-architecture/domain/models"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/responses"
	"clean-architecture/pkg/utils"
	"net/http"
	"strconv"

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

	if err := c.ShouldBindJSON(&user); err != nil {
		responses.HandleValidationError(u.logger, c, err)
		return
	}

	// check if the user already exists
	exists, err := u.service.ExistsByEmail(user.Email)
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}
	if exists {
		responses.Error(c, ErrUserAlreadyExists.WithDetail("email", user.Email))
		return
	}

	if err := u.service.Create(&user); err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	responses.Success(c, http.StatusCreated, user, nil)
}

// GetOneUser gets one user
func (u *Controller) GetUserByID(c *gin.Context) {

	paramID := c.Param("id")
	uid, err := strconv.ParseUint(paramID, 10, 64)
	if err != nil {
		responses.HandleValidationError(u.logger, c, ErrInvalidUserID)
		return
	}
	user, err := u.service.GetUserByID(uint(uid))
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})

}

func (u *Controller) GetUsers(c *gin.Context) {
	u.logger.Info("[UserController...GetUsers]")
	pagination := utils.BuildPagination(c)
	users, count, err := u.service.GetUsers(pagination)
	if err != nil {
		responses.HandleError(u.logger, c, err)
		return
	}
	responses.PaginationSuccess(c, http.StatusOK, users, count)

}
