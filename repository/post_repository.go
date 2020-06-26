package repository

import (
	"prototype2/entity"
	"fmt"

	// "github.com/jinzhu/gorm"
  "github.com/gin-gonic/gin"
)

type PostRepository struct {
	c *gin.Context
}

//New repository constructor
func NewPostRepository() DatabaseRepository {
	return &PostRepository{}
}


func (p *PostRepository) Save(post *entity.Post) (*entity.Post, error) {
	return nil, nil
}

func (p *PostRepository) FindAll() ([]entity.Post, error) {
	fmt.Println("Lalalalalalalalala")
	fmt.Println(p.c)
	return nil, nil
}

func (p *PostRepository) FindByID(ID int64) (*entity.Post, error) {
	return nil, nil
}