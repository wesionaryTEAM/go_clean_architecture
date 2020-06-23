package controller

import (
	"encoding/json"
	"net/http"
	"math/rand"
	"prototype2/entity"
	repository "prototype2/repository"
	service "prototype2/service"
)

type PostController interface {
	GetPosts(response http.ResponseWriter, request *http.Request)
	AddPost(response http.ResponseWriter, request (http.Request))
}

var (
	// repo repository.PostRepository = repository.NewPostRepository()
	postService service.PostService = service.NewPostService()
)

func GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": Error getting the posts"`))
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")	
	var post entity.Post 
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error marshalling the request"`))
		return
	}

	post.ID = rand.Int63()

	err := postService.Validate(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Validation error"`))
		return
	}
	post, err := postService.Create(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`"error": "Error while creating post"`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}