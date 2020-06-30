package service

import (
	"prototype2/domain"
	"prototype2/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidateEmptyPostField(t *testing.T) {
	testService := NewPostService(nil)

	scenarios := []struct {
		post   domain.Post
		expect string
	}{
		{post: domain.Post{ID: 1, Title: "", Text: "Ball"}, expect: "The post title is empty"},
		{post: domain.Post{ID: 1, Title: "Apple", Text: ""}, expect: "The post text is empty"},
	}

	for _, s := range scenarios {
		err := testService.Validate(&s.post)
		assert.NotNil(t, err)
		assert.Equal(t, s.expect, err.Error())
	}
}

func TestFindAll(t *testing.T) {
	mockRepo := new(mocks.PostRepository)

	var identifier int64 = 1

	post := domain.Post{ID: identifier, Title: "Apple", Text: "Ball"}

	mockRepo.On("FindAll").Return([]domain.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "Apple", result[0].Title)
	assert.Equal(t, "Ball", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(mocks.PostRepository)
	post := domain.Post{Title: "Test Title", Text: "Test Text"}

	//Setting up the expectations
	mockRepo.On("Save", &post).Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "Test Title", result.Title)
	assert.Equal(t, "Test Text", result.Text)
	assert.Nil(t, err)
}
