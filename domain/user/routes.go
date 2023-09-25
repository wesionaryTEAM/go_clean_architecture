package user

import (
	"clean-architecture/domain/constants"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/middlewares"
)

// UserRoutes struct
type Routes struct {
	logger           framework.Logger
	handler          infrastructure.Router
	controller       *Controller
	authMiddleware   middlewares.AuthMiddleware
	uploadMiddleware middlewares.UploadMiddleware
	middlewares.PaginationMiddleware
	rateLimitMiddleware middlewares.RateLimitMiddleware
}

func NewRoutes(
	logger framework.Logger,
	handler infrastructure.Router,
	controller *Controller,
	authMiddleware middlewares.FirebaseAuthMiddleware,
	uploadMiddleware middlewares.UploadMiddleware,
	pagination middlewares.PaginationMiddleware,
	rateLimit middlewares.RateLimitMiddleware,
) {
	r := &Routes{
		handler:              handler,
		logger:               logger,
		controller:           controller,
		authMiddleware:       authMiddleware,
		uploadMiddleware:     uploadMiddleware,
		PaginationMiddleware: pagination,
		rateLimitMiddleware:  rateLimit,
	}
	r.Setup()
}

// Setup user routes
func (r *Routes) Setup() {
	r.logger.Info("Setting up routes")

	// in HandleAuthWithRole() pass empty for authentication
	// or pass user role for authentication along with authorization
	api := r.handler.Group("/api").Use(r.authMiddleware.HandleAuthWithRole(constants.RoleIsAdmin),
		r.rateLimitMiddleware.Handle())

	api.GET("/user", r.PaginationMiddleware.Handle(), r.controller.GetUser)
	api.GET("/user/:id", r.controller.GetOneUser)
	api.POST("/user", r.controller.SaveUser)
	api.PUT("/user/:id",
		r.uploadMiddleware.Push(r.uploadMiddleware.Config().ThumbEnable(true).WebpEnable(true)).Handle(),
		r.controller.UpdateUser,
	)
	api.DELETE("/user/:id", r.controller.DeleteUser)

}
