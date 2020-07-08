package controller

import (
	"math/rand"
	"net/http"
	"prototype2/domain"
	// "prototype2/service/post"

	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
)

type postController struct {
	postService domain.PostService
}

// PostController : set of methods in post controller
type PostController interface {
	GetPosts(c *gin.Context)
	AddPost(c *gin.Context)
}

// NewPostController : get injected post service
func NewPostController(s domain.PostService) PostController {
	return &postController{
		postService: s,
	}
}

func (p *postController) GetPosts(c *gin.Context) {
	posts, err := p.postService.FindAll()
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error getting the posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (p *postController) AddPost(c *gin.Context) {
	var post domain.Post
	if err := c.ShouldBindJSON(&post); err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post.ID = rand.Int63()

	if err1 := p.postService.Validate(&post); err1 != nil {
		sentry.CaptureException(err1)
		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}
	p.postService.Create(&post)
	c.JSON(http.StatusOK, post)
}

/** Implementation of the controller for the MUX or CHI Router but not GIN router */
/**
* Implementation only for the controller level testing
 */

// import (
// 	// "math/rand"
// 	"net/http"
// 	"prototype2/domain"
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
// 	var post domain.Post
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
