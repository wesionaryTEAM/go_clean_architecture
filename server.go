package main

import (
	"log"
	"net/http"
	"prototype2/controller"
	router "prototype2/http"
	"prototype2/infrastructure"
	"prototype2/repository"
	"prototype2/service"
	"prototype2/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	postRepository repository.PostRepository
	postService    service.PostService
	postController controller.PostController
	httpRouter     router.Router = router.NewGinRouter()
)

func main() {
	utils.LoadEnv()

	db := infrastructure.SetupModels()

	postRepository = repository.NewPostRepository(db)
	if err := postRepository.Migrate(); err != nil {
		log.Fatal("post migrate err", err)
	}

	postService = service.NewPostService(postRepository)

	postController = controller.NewPostController(postService)

	const port string = ":8000"

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
