package service

import (
	"fmt"
	"math/rand"
	"prototype2/entity"
	repository "prototype2/repository"
)

type PostService interface {
	Validate(post *entity.Post) error 
	Create(post  *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct {}

var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

func NewPostService() PostService{	
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if(post == nil){
		fmt.Println("The post is empty")
		// err := errors.New("The post is empty")
		// return err
	}
	if(post.Title == ""){
		// err := errors.New("The post title is empty")
		// return err
		fmt.Println("The post is empty")
	}
	return nil
}

func (*service) Create(post  *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
}