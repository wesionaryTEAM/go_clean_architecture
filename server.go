package main

import (
	"fmt"
	"log"
	"net/http"
	"prototype2/controller"
	"prototype2/service"
	"prototype2/repository"
	"prototype2/http"
	"prototype2/config"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	postRepository repository.DatabaseRepository = repository.NewPostRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter router.Router = router.NewGinRouter()
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatalf("Error getting env, %v", err)
  } else {
    fmt.Println("Environment variables fetched ")
	}
	
	//Init database
	var db = config.InitDatabase()
	var gc = func(c *gin.Context){
		c.Set("db", db)
	}
	
	fmt.Println(gc)

	const port string = ":8000"

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	// httpRouter.GET("/users", repository.GetUsers)
	// httpRouter.POST("/users", repository.AddUser)
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}