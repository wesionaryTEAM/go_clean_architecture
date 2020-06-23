package main

import (
	"fmt"
	"net/http"
	router "prototype2/http"
	"prototype2/controller"
)

var (
	postController controller.PostController = controller.NewPostController()
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running....")
	})
	httpRouter.GET("/posts", postController.GetPosts).Methods("GET")
	httpRouter.POST("/posts", postController.AddPost).Methods("POST")
	httpRouter.SERVE(port)
}