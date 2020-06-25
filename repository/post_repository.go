package repository

import (
	"prototype2/entity"
)

type PostRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Migrate() error
}
