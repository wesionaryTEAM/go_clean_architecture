package routes

import (
	"clean-architecture/api/controllers"
	"clean-architecture/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        lib.Router
	userController controllers.UserController
}

func NewUserRoutes(logger lib.Logger, handler lib.Router, userController controllers.UserController) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
	}
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Gin.Group("/users")
	{
		api.GET("", s.userController.GetUsers)
	}
}
