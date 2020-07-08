package service

import (
	"errors"
	"math/rand"
	"prototype2/domain"
	"sync"
)

var once sync.Once

type postService struct {
	repo domain.PostRepository
}

var instance *postService

// NewPostService : get injected post repository
func NewPostService(r domain.PostRepository) domain.PostService {
	once.Do(func() {
		instance = &postService{
			repo: r,
		}
	})
	return instance
}

func (*postService) Validate(post *domain.Post) error {
	if post == nil {
		err := errors.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The post title is empty")
		return err
	}
	if post.Text == "" {
		err := errors.New("The post text is empty")
		return err
	}
	return nil
}

func (p *postService) Create(post *domain.Post) (*domain.Post, error) {
	post.ID = rand.Int63()
	return p.repo.Save(post)
}

func (p *postService) FindAll() ([]domain.Post, error) {
	return p.repo.FindAll()
}
