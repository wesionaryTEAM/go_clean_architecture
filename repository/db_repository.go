package repository

import (
	"prototype2/entity"

	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Delete(post *entity.Post) error	
	Migrate() error
}