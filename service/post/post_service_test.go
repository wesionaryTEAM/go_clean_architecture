package service

import (
	"prototype2/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)

	assert.Equal(t, "The post is empty", err.Error())
}

func TestValidateEmptyPostField(t *testing.T) {
	testService := NewPostService(nil)

	scenarios := []struct {
		post   entity.Post
		expect string
	}{
		{post: entity.Post{ID: 1, Title: "", Text: "Ball"}, expect: "The post title is empty"},
		{post: entity.Post{ID: 1, Title: "Apple", Text: ""}, expect: "The post text is empty"},
	}

	for _, s := range scenarios {
		err := testService.Validate(&s.post)
		assert.NotNil(t, err)
		assert.Equal(t, s.expect, err.Error())
	}
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier int64 = 1

	post := entity.Post{ID: identifier, Title: "Apple", Text: "Ball"}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "Apple", result[0].Title)
	assert.Equal(t, "Ball", result[0].Text)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)
	post := entity.Post{Title: "Test Title", Text: "Test Text"}

	//Setting up the expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.NotNil(t, result.ID)
	assert.Equal(t, "Test Title", result.Title)
	assert.Equal(t, "Test Text", result.Text)
	assert.Nil(t, err)
}
