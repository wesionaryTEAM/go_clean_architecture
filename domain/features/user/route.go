package user

import (
	"clean-architecture/domain/constants"
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
	"clean-architecture/pkg/interfaces"
	"clean-architecture/pkg/middlewares"
)

// UserRoutes struct
type Route struct {
	logger     framework.Logger
	handler    infrastructure.Router
	controller *Controller
	interfaces.PaginationMiddleware
	rateLimitMiddleware middlewares.RateLimitMiddleware
	authMiddleware      interfaces.AuthMiddleware
}

func NewRoute(
	logger framework.Logger,
	handler infrastructure.Router,
	controller *Controller,
	pagination interfaces.PaginationMiddleware,
	rateLimit middlewares.RateLimitMiddleware,
	authMiddleware interfaces.AuthMiddleware,
) *Route {
	return &Route{
		handler:              handler,
		logger:               logger,
		controller:           controller,
		PaginationMiddleware: pagination,
		rateLimitMiddleware:  rateLimit,
		authMiddleware:       authMiddleware,
	}

}

// Setup user routes
func RegisterRoute(r *Route) {
	r.logger.Info("Setting up routes")

	// in HandleAuthWithRole() pass empty for authentication
	// or pass user role for authentication along with authorization
	api := r.handler.Group("/api").Use(r.authMiddleware.HandleAuthWithRole(constants.RoleIsAdmin))
	// 	r.rateLimitMiddleware.Handle())

	api.GET("/user",
		r.PaginationMiddleware.Handle(),
		r.controller.GetUser)
	api.GET("/user/:id", r.controller.GetOneUser)
	api.POST("/user", r.controller.SaveUser)
	// api.PUT("/user/:id",
	// 	r.uploadMiddleware.Push(r.uploadMiddleware.Config().ThumbEnable(true).WebpEnable(true)).Handle(),
	// 	r.controller.UpdateUser,
	// )
	api.DELETE("/user/:id", r.controller.DeleteUser)

}
