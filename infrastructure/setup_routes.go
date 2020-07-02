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

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//SetupRoutes : all the routes are defined Here
func SetupRoutes(db *gorm.DB) {

	httpRouter := router.NewGinRouter()

	r := gin.Default()

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

	// dependency injection for user resources
	userRepository := user_repository.NewUserRepository(db)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("user migrate err", err)
	}
	userService := user_service.NewUserService(userRepository)
	userController := user_controller.NewUserController(userService)

	/* Middleware checks */
	userAuthentication := r.Group("/")

	/* Routes */
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	userAuthentication.Use(middleware.UserAuthenticationMiddleware)
	{
		httpRouter.GET("/users", userController.GetUsers)
		httpRouter.POST("/users", userController.AddUser)
	}

	httpRouter.SERVE(port)
}
