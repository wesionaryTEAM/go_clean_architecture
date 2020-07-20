package controller

// import (
// 	"bytes"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"encoding/json"
// 	"prototype2/repository"
// 	"io"
// 	"prototype2/service"
// 	"prototype2/domain"
// 	"prototype2/infrastructure"

// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	db = infrastructure.SetupModelsForControllerTest()
// 	postRepository = repository.NewPostRepository(db)
// 	postService service.PostService = service.NewPostService(postRepository)
// 	postController PostController = NewPostController(postService)
// )

// const (
// 	ID int64 = 123
// 	TITLE string = "Lorem Ipsum"
// 	TEXT string = "Lorem Ipsum Lorem Ipsum Lorem Ipsum"
// )

// func TestAddPost(t *testing.T){
// 	//Create a new HTTP POST request
// 	var jsonString = []byte(`{"title":"` + TITLE + `","text":"` + TEXT + `"}`)
// 	request, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonString))

// 	//assign the controller action as a handler function (add post action)
// 	handler := http.HandlerFunc(postController.AddPost)

// 	//Record the http response using httptest
// 	response := httptest.NewRecorder()

// 	// dispatch http request
// 	handler.ServeHTTP(response, request)

// 	// add assertions on the http status code and the response
// 	status := response.Code

// 	if status != http.StatusOK {
// 		t.Errorf("Handler returned wrong status code: got %v want %v",
// 		status, http.StatusOK)
// 	}

// 	//Decode the HTTP response
// 	var post domain.Post
// 	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

// 	//Assert HTTP response
// 	assert.NotNil(t, post.ID)
// 	assert.Equal(t, TITLE, post.Title)
// 	assert.Equal(t, TEXT, post.Text)

// 	// Cleanup database
// 	cleanUp(&post)
// }

// func cleanUp(post *domain.Post){
// 	postRepository.Delete(post)
// }

// func setup() (*domain.Post, error) {
// 	var post domain.Post = domain.Post{
// 		ID:    ID,
// 		Title: TITLE,
// 		Text:  TEXT,
// 	}
// 	return postRepository.Save(&post)
// }

// func TestGetPosts(t *testing.T) {

// 	// Insert new post
// 	post, err := setup()
// 	if err != nil {
// 		t.Errorf("Error occurred initializing the post")
// 	}

// 	// Create new HTTP request
// 	req, _ := http.NewRequest("GET", "/posts", nil)

// 	// Assing HTTP Request handler Function (controller function)
// 	handler := http.HandlerFunc(postController.GetPosts)
// 	// Record the HTTP Response
// 	response := httptest.NewRecorder()
// 	// Dispatch the HTTP Request
// 	handler.ServeHTTP(response, req)

// 	// Assert HTTP status
// 	status := response.Code
// 	if status != http.StatusOK {
// 		t.Errorf("Handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}

// 	// Decode HTTP response
// 	var posts []domain.Post
// 	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

// 	// Assert HTTP response
// 	assert.Equal(t, ID, posts[0].ID)
// 	assert.Equal(t, TITLE, posts[0].Title)
// 	assert.Equal(t, TEXT, posts[0].Text)

// 	// Cleanup database
// 	cleanUp(post)
// }
