package main

import (
	"fmt"
	"log"
	"net/http"
	// "encoding/json"
	// "math/rand"
	"prototype2/routes"

	// "./entity"
	// "./repository"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter();
	const port string = ":8000"
	router.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running....")
	})
	router.HandleFunc("/posts", routes.GetPosts).Methods("GET")
	router.HandleFunc("/posts", routes.AddPost).Methods("POST")
	log.Print("server listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router)) 

}