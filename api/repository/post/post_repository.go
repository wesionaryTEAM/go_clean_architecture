package repository

import (
	"log"
	"prototype2/domain"

	"github.com/jinzhu/gorm"
)

type postRepository struct {
	DB *gorm.DB
}

// NewPostRepository : get injected database
func NewPostRepository(db *gorm.DB) domain.PostRepository {
	return &postRepository{
		DB: db,
	}
}

func (p *postRepository) Save(post *domain.Post) (*domain.Post, error) {
	log.Print("[PostRepository]...Save")
	return post, p.DB.Create(post).Error
}

func (p *postRepository) FindAll() ([]domain.Post, error) {
	log.Print("[PostRepository]...FindAll")
	var posts []domain.Post
	err := p.DB.Find(&posts).Error
	return posts, err
}

func (p *postRepository) Delete(post *domain.Post) error {
	log.Print("[PostRepository]...Delete")
	return p.DB.Delete(&post).Error
}

func (p *postRepository) Migrate() error {
	log.Print("[PostRepository]...Migrate")
	return p.DB.AutoMigrate(&domain.Post{}).Error
}
