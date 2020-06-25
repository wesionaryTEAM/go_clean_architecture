package repository

import (
	"prototype2/entity"

	"github.com/jinzhu/gorm"
)

type repo struct {
	DB *gorm.DB
}

func NewSqlRepository(db *gorm.DB) PostRepository {
	return &repo{DB: db}
}

func (r *repo) Save(post *entity.Post) (*entity.Post, error) {
	return post, r.DB.Create(post).Error
}

func (r *repo) FindAll() ([]entity.Post, error) {
	var posts []entity.Post
	err := r.DB.Find(&posts).Error
	return posts, err
}

func (r *repo) Migrate() error {
	return r.DB.AutoMigrate(&entity.Post{}).Error
}
