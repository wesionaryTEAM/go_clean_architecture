package main

import (
	"log"
	"net/http"
	postcontroller "prototype2/controller/post"
	usercontroller "prototype2/controller/user"
	"prototype2/domain"
	router "prototype2/http"
	"prototype2/infrastructure"
	postrepository "prototype2/repository/post"
	userrepository "prototype2/repository/user"
	postservice "prototype2/service/post"
	userservice "prototype2/service/user"
	"prototype2/utils"

	"github.com/gin-gonic/gin"
)

var (
	postRepository domain.PostRepository
	postService    domain.PostService
	postController postcontroller.PostController

	userRepository domain.UserRepository
	userService		 domain.UserService
	userController usercontroller.UserController
	
	httpRouter     router.Router = router.NewGinRouter()
)

func main() {
	utils.LoadEnv()

	db := infrastructure.SetupModels()

	postRepository = postrepository.NewPostRepository(db)
	if err := postRepository.Migrate(); err != nil {
		log.Fatal("post migrate err", err)
	}

	postService = postservice.NewPostService(postRepository)

	postController = postcontroller.NewPostController(postService)

	userRepository = userrepository.NewUserRepository(db)
	if err := userRepository.Migrate(); err != nil {
		log.Fatal("user migrate err", err)
	}

	userService = userservice.NewUserService(userRepository)

	userController = usercontroller.NewUserController(userService)

	const port string = ":8000"

	httpRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Up and Running..."})
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/post", postController.AddPost)
	httpRouter.GET("/users", userController.GetUsers)
	httpRouter.POST("/users", userController.AddUser)
	httpRouter.SERVE(port)
}
