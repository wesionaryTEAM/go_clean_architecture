package user

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/middlewares"
)

// UserRoutes struct
type Route struct {
	logger           framework.Logger
	handler          infrastructure.Router
	controller       *Controller
	authMiddleware   middlewares.CognitoMiddleWare
	uploadMiddleware middlewares.UploadMiddleware
	middlewares.PaginationMiddleware
	rateLimitMiddleware middlewares.RateLimitMiddleware
}

func NewRoute(
	logger framework.Logger,
	handler infrastructure.Router,
	controller *Controller,
	authMiddleware middlewares.CognitoAuthMiddleware,
	uploadMiddleware middlewares.UploadMiddleware,
	pagination middlewares.PaginationMiddleware,
	rateLimit middlewares.RateLimitMiddleware,
) *Route {
	return &Route{
		handler:              handler,
		logger:               logger,
		controller:           controller,
		authMiddleware:       authMiddleware,
		uploadMiddleware:     uploadMiddleware,
		PaginationMiddleware: pagination,
		rateLimitMiddleware:  rateLimit,
	}

}

// Setup user routes
func RegisterRoute(r *Route) {
	r.logger.Info("Setting up routes")

	// remove r.authMiddleware.Handle() if you don't have the access token
	api := r.handler.Group("/api").Use(r.authMiddleware.Handle(), r.rateLimitMiddleware.Handle())

	api.GET("/user", r.PaginationMiddleware.Handle(), r.controller.GetUser)
	api.GET("/user/:id", r.controller.GetOneUser)
	api.POST("/user", r.controller.SaveUser)
	api.PUT("/user/:id",
		r.uploadMiddleware.Push(r.uploadMiddleware.Config().ThumbEnable(true).WebpEnable(true)).Handle(),
		r.controller.UpdateUser,
	)
	api.DELETE("/user/:id", r.controller.DeleteUser)
	api.POST("/upload-test", r.uploadMiddleware.Push(r.uploadMiddleware.Config().ThumbEnable(true).WebpEnable(true)).Handle(), r.controller.UploadImage)
	api.POST("/send-test-email", r.controller.SendEmail)
}
