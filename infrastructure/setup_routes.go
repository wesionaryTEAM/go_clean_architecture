package infrastructure

import (
	"log"
	"net/http"
	// post controller, repository and service
	post_controller "prototype2/controller/post"
	post_repository "prototype2/repository/post"
	post_service "prototype2/service/post"
	// user controlelr, repository and service
	user_controller "prototype2/controller/user"
	user_repository "prototype2/repository/user"
	user_service "prototype2/service/user"
	// Add other controllers, repository and services below
	router "prototype2/http"
	"prototype2/middleware"
	"prototype2/service"

	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//SetupRoutes : all the routes are defined Here
func SetupRoutes(db *gorm.DB, fb *auth.Client) {

	httpRouter := router.NewGinRouter()

	const port string = ":8000"

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	// dependency injections for post resources
	postRepository := post_repository.NewPostRepository(db)
	if err := postRepository.Migrate(); err != nil {
		log.Fatal("post migrate err", err)
	}
	postService := post_service.NewPostService(postRepository)
	postController := post_controller.NewPostController(postService)

	// Post routes
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	// dependency injection for firebase middleware
	fbAuth := service.NewFirebaseService(fb)
	auth := middleware.NewMiddlewareAuth(fbAuth, "user")

	// dependency injection for user resources
	userRepository := user_repository.NewUserRepository(db)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("user migrate err", err)
	}
	userService := user_service.NewUserService(userRepository)
	userController := user_controller.NewUserController(userService, fbAuth)

	// User Routes
	httpRouter.POST("/users", userController.AddUser)
	users := httpRouter.GROUP("/")
	users.Use(auth)
	{
		users.GET("/users", userController.GetUsers)
	}

	// Run server
	httpRouter.SERVE(port)
}
