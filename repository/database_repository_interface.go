package repository

import (
	"prototype2/entity"
)

type DatabaseRepository interface {
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	FindByID(ID int64) (*entity.Post, error)
}