package main

import (
	"net/http"
	"prototype2/controller"
	"prototype2/service"
	"prototype2/repository"
	"prototype2/http"

	"github.com/gin-gonic/gin"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter router.Router = router.NewGinRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	httpRouter.GET("/users", repository.GetUsers)
	httpRouter.POST("/users", repository.AddUser)
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
