package routes

import (
	"log"

	// post controller, repository and service
	post_controller "prototype2/api/controller/post"
	post_repository "prototype2/api/repository/post"
	post_service "prototype2/api/service/post"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// PostRoutes - all post routes are placed here "/posts/***"
func PostRoutes(route *gin.RouterGroup, db *gorm.DB) {
	// dependency injections for post resources
	postRepository := post_repository.NewPostRepository(db)
	if err := postRepository.Migrate(); err != nil {
		sentry.CaptureException(err)
		log.Fatal("post migrate err", err)
	}
	postService := post_service.NewPostService(postRepository)
	postController := post_controller.NewPostController(postService)

	// Post routes
	route.GET("/", postController.GetPosts)
	route.POST("/", postController.AddPost)
	route.GET("/:id", postController.GetPost)
	route.DELETE("/:id", postController.DeletePost)
}
