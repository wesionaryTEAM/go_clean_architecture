package routes

import (
	"clean-architecture/api/controllers"
	"clean-architecture/api/middlewares"
	"clean-architecture/constants"
	"clean-architecture/infrastructure"
	"clean-architecture/lib"
)

// UserRoutes struct
type UserRoutes struct {
	logger           lib.Logger
	handler          infrastructure.Router
	userController   *controllers.UserController
	authMiddleware   middlewares.FirebaseAuthMiddleware
	uploadMiddleware middlewares.UploadMiddleware
	middlewares.PaginationMiddleware
	rateLimitMiddleware middlewares.RateLimitMiddleware
}

func NewUserRoutes(
	logger lib.Logger,
	handler infrastructure.Router,
	userController *controllers.UserController,
	authMiddleware middlewares.FirebaseAuthMiddleware,
	uploadMiddleware middlewares.UploadMiddleware,
	pagination middlewares.PaginationMiddleware,
	rateLimit middlewares.RateLimitMiddleware,
) *UserRoutes {
	return &UserRoutes{
		handler:              handler,
		logger:               logger,
		userController:       userController,
		authMiddleware:       authMiddleware,
		uploadMiddleware:     uploadMiddleware,
		PaginationMiddleware: pagination,
		rateLimitMiddleware:  rateLimit,
	}
}

// Setup user routes
func (s *UserRoutes) Setup() {
	s.logger.Info("Setting up routes")

	api := s.handler.Group("/api").Use(s.authMiddleware.HandleAuthWithRole(constants.RoleIsAdmin),
		s.rateLimitMiddleware.Handle())

	api.GET("/user", s.PaginationMiddleware.Handle(), s.userController.GetUser)
	api.GET("/user/:id", s.userController.GetOneUser)
	api.POST("/user", s.userController.SaveUser)
	api.PUT("/user/:id",
		s.uploadMiddleware.Push(s.uploadMiddleware.Config().ThumbEnable(true).WebpEnable(true)).Handle(),
		s.userController.UpdateUser,
	)
	api.DELETE("/user/:id", s.userController.DeleteUser)

}
