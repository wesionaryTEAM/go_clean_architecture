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
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetPosts(c *gin.Context)
	AddPost(c *gin.Context)
}

func NewPostController() PostController {
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
