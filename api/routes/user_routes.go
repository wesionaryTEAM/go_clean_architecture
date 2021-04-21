package routes

import (
	"clean-architecture/api/controllers"
	"clean-architecture/api/middlewares"
	"clean-architecture/infrastructure"
)

// UserRoutes struct
type UserRoutes struct {
	logger         infrastructure.Logger
	handler        infrastructure.Router
	userController controllers.UserController
	authMiddleware middlewares.FirebaseAuthMiddleware
}

func NewUserRoutes(
	logger infrastructure.Logger,
	handler infrastructure.Router,
	userController controllers.UserController,
	authMiddleware middlewares.FirebaseAuthMiddleware,
) UserRoutes {
	return UserRoutes{
		handler:        handler,
		logger:         logger,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}

// Setup user routes
func (s UserRoutes) Setup() {
	s.logger.Zap.Info("Setting up routes")
	api := s.handler.Group("/api").Use(s.authMiddleware.Handle())
	{
		api.GET("/user", s.userController.GetUser)
		api.GET("/user/:id", s.userController.GetOneUser)
		api.POST("/user", s.userController.SaveUser)
		api.POST("/user/:id", s.userController.UpdateUser)
		api.DELETE("/user/:id", s.userController.DeleteUser)
	}
}
