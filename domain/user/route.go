package user

import (
	"clean-architecture/pkg/framework"
	"clean-architecture/pkg/infrastructure"
)

// UserRoutes struct
type Route struct {
	logger     framework.Logger
	handler    infrastructure.Router
	controller *Controller
}

func NewRoute(
	logger framework.Logger,
	handler infrastructure.Router,
	controller *Controller,
) *Route {
	return &Route{
		handler:    handler,
		logger:     logger,
		controller: controller,
	}

}

// Setup user routes
func RegisterRoute(r *Route) {
	r.logger.Info("Setting up routes")

	api := r.handler.Group("/api")

	api.POST("/user", r.controller.CreateUser)
	api.GET("/user/:id", r.controller.GetUserByID)
	api.GET("/users", r.controller.GetUsers)

}
