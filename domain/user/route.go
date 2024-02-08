package user

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/middlewares"
)

// UserRoutes struct
type Route struct {
	logger     framework.Logger
	handler    infrastructure.Router
	controller *Controller
	middlewares.PaginationMiddleware
	rateLimitMiddleware middlewares.RateLimitMiddleware
}

func NewRoute(
	logger framework.Logger,
	handler infrastructure.Router,
	controller *Controller,
	pagination middlewares.PaginationMiddleware,
	rateLimit middlewares.RateLimitMiddleware,
) *Route {
	return &Route{
		handler:              handler,
		logger:               logger,
		controller:           controller,
		PaginationMiddleware: pagination,
		rateLimitMiddleware:  rateLimit,
	}

}

// Setup user routes
func RegisterRoute(r *Route) {
	r.logger.Info("Setting up routes")

	// in HandleAuthWithRole() pass empty for authentication
	// or pass user role for authentication along with authorization
	api := r.handler.Group("/api")
	// .Use(r.authMiddleware.HandleAuthWithRole(constants.RoleIsAdmin),
	// 	r.rateLimitMiddleware.Handle())

	api.GET("/user", r.PaginationMiddleware.Handle(), r.controller.GetUser)
	api.GET("/user/:id", r.controller.GetOneUser)
	api.POST("/user", r.controller.SaveUser)
	// api.PUT("/user/:id",
	// 	r.uploadMiddleware.Push(r.uploadMiddleware.Config().ThumbEnable(true).WebpEnable(true)).Handle(),
	// 	r.controller.UpdateUser,
	// )
	api.DELETE("/user/:id", r.controller.DeleteUser)

}
