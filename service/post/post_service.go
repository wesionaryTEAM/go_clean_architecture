package service

import (
	"errors"
	"math/rand"
	"prototype2/domain"
)

type postService struct {
	repo domain.PostRepository
}

// NewPostService : get injected post repository
func NewPostService(r domain.PostRepository) domain.PostService {
	return &postService{
		repo: r,
	}
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
