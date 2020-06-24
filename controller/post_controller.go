package controller

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"prototype2/entity"
	service "prototype2/service"
)

type controller struct {}

var (
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request *http.Request)
}

func NewPostController() PostController {
	return &controller{}
}

func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`"error": Error getting the posts"`))
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")	
	var post entity.Post 
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`"error": "Error marshalling the request"`))
		return
	}

	post.ID = rand.Int63()

	err = postService.Validate(&post)
	postService.Create(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)
}