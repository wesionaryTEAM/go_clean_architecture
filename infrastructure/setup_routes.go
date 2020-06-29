package infrastructure

import (
	"log"
	"net/http"
	postCon "prototype2/controller/post"
	router "prototype2/http"
	postRepo "prototype2/repository/post"
	postServ "prototype2/service/post"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

//SetupRoutes : all the routes are defined Here
func SetupRoutes(db *gorm.DB) {

	httpRouter := router.NewGinRouter()

	const port string = ":8000"

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	postRepository := postRepo.NewPostRepository(db)
	if err := postRepository.Migrate(); err != nil {
		log.Fatal("post migrate err", err)
	}
	postService := postServ.NewPostService(postRepository)
	postController := postCon.NewPostController(postService)

	// post routes
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/post", postController.AddPost)

	httpRouter.SERVE(port)
}
