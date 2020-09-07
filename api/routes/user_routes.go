package routes

import (
	"log"

	// user controller, repository and service
	user_controller "prototype2/api/controller/user"
	user_repository "prototype2/api/repository/user"
	user_service "prototype2/api/service/user"

	"prototype2/api/middleware"
	"prototype2/api/service"

	"firebase.google.com/go/auth"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// UserRoutes - all user routes are placed here "/users/***"
func UserRoutes(route *gin.RouterGroup, db *gorm.DB, fb *auth.Client) {
	fbAuth := service.NewFirebaseService(fb)

	// dependency injection for user resources
	userRepository := user_repository.NewUserRepository(db)
	if err := userRepository.Migrate(); err != nil {
		sentry.CaptureException(err)
		log.Fatal("user migrate err", err)
	}
	userService := user_service.NewUserService(userRepository)
	userController := user_controller.NewUserController(userService, fbAuth)

	// Initialize MiddlewareAuth for users
	userAuth := middleware.NewMiddlewareAuth(fbAuth, "user")

	// User Routes
	route.POST("/", userController.AddUser)
	users := route.Group("/")
	users.Use(userAuth)
	{
		users.GET("/", userController.GetUsers)
	}
}
