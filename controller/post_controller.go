package controller

import (
	"math/rand"
	"net/http"
	"prototype2/entity"
	service "prototype2/service"

	"github.com/gin-gonic/gin"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(c *gin.Context)
	AddPost(c *gin.Context)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(c *gin.Context) {
	posts, err := postService.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting the posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (*controller) AddPost(c *gin.Context) {
	var post entity.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.ID = rand.Int63()

	if err1 := postService.Validate(&post); err1 != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	postService.Create(&post)
	c.JSON(http.StatusOK, post)
}



/** Implementation of the controller for the MUX or CHI Router but not GIN router */
/**
* Implementation only for the controller level testing
*/

// import (
// 	// "math/rand"
// 	"net/http"
// 	"prototype2/entity"
// 	service "prototype2/service"
// 	"encoding/json"
// 	"errors"

// 	// "github.com/gin-gonic/gin"
// )

// type controller struct{}

// type PostController interface {
// 	GetPosts(response http.ResponseWriter, request *http.Request)
// 	AddPost(response http.ResponseWriter, request *http.Request)
// }

// func NewPostController(service service.PostService) PostController {
// 	postService = service
// 	return &controller{}
// }

// func (*controller) GetPosts(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("Content-Type", "application/json")
// 	posts, err := postService.FindAll()
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(response).Encode(errors.New("Error getting the posts"))
// 	}
// 	response.WriteHeader(http.StatusOK)
// 	json.NewEncoder(response).Encode(posts)
// }

// func (*controller) AddPost(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("Content-Type", "application/json")
// 	var post entity.Post
// 	err := json.NewDecoder(request.Body).Decode(&post)
// 	if err != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(response).Encode(errors.New("Error unmarshalling data"))
// 		return
// 	}
// 	err1 := postService.Validate(&post)
// 	if err1 != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(response).Encode(errors.New("Error while adding the post"))
// 		return
// 	}

// 	result, err2 := postService.Create(&post)
// 	if err2 != nil {
// 		response.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(response).Encode(errors.New("Error saving the post"))
// 		return
// 	}
// 	response.WriteHeader(http.StatusOK)
// 	json.NewEncoder(response).Encode(result)
// }
