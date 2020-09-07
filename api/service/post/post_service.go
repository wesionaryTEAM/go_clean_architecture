package service

import (
	"log"
	"math/rand"
	"strconv"
	"sync"

	"prototype2/domain"
	"prototype2/errors"
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
	// return &postService{
	// 	repo: r,
	// }
}

func (*postService) Validate(post *domain.Post) error {
	log.Print("[PostService]...Validate")
	if post == nil {
		err := errors.BadRequest.New("The post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.BadRequest.New("The post title is empty")
		return err
	}
	if post.Text == "" {
		err := errors.BadRequest.New("The post text is empty")
		return err
	}
	return nil
}

func (p *postService) Create(post *domain.Post) (*domain.Post, error) {
	log.Print("[PostService]...Create")
	post.ID = rand.Int63()
	return p.repo.Save(post)
}

func (p *postService) FindAll() ([]domain.Post, error) {
	log.Print("[PostService]...FindAll")
	return p.repo.FindAll()
}

func (p *postService) GetByID(idString string) (*domain.Post, error) {
	log.Print("[PostService]...GetByID")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		err = errors.BadRequest.Wrapf(err, "interactor converting id to int")
		err = errors.AddErrorContext(err, "id", "wrong id format")

		return nil, err
	}

	return p.repo.FindByID(id)
}

func (p *postService) Delete(idString string) error {
	log.Print("[PostService]...Delete")

	id, err := strconv.ParseInt(idString, 10, 64)
	if err != nil {
		err = errors.BadRequest.Wrapf(err, "interactor converting id to int")
		err = errors.AddErrorContext(err, "id", "wrong id format")

		return err
	}

	post, err := p.repo.FindByID(id)
	if err != nil {
		return err
	}
	return p.repo.Delete(post)
}
