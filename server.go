package main

import (
	"fmt"
	"log"
	// "net/http"
	// "encoding/json"
	// "math/rand"
	// "prototype2/routes"
	router "prototype2/http"

	// "./entity"
	// "./repository"
	// "github.com/gorilla/mux"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running....")
	})
	httpRouter.GET("/posts", routes.GetPosts).Methods("GET")
	httpRouter.POST("/posts", routes.AddPost).Methods("POST")
	httpRouter.SERVE(port)

}