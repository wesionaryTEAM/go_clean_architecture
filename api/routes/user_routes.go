package routes

import (
	"prototype2/api/handlers"
	"prototype2/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	userController handlers.UserController
}


func NewUserRoutes(logger lib.Logger, handler lib.RequestHandler, userController handlers.UserController) UserRoutes {
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
