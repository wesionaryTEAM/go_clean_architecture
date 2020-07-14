package repository

import (
	"log"
	"prototype2/domain"
	"prototype2/errors"

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
	result := p.DB.Find(&posts)

	switch result.Error {
	case nil:
		return posts, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.ErrNotFound
	default:
		return nil, errors.ErrDatabase
	}
}

func (p *postRepository) FindByID(id int64) (*domain.Post, error) {
	log.Print("[PostRepository]...FindAll")
	var post domain.Post
	result := p.DB.Where("id = ?", id).First(&post)

	switch result.Error {
	case nil:
		return &post, nil
	case gorm.ErrRecordNotFound:
		return nil, errors.ErrNotFound
	default:
		return nil, errors.ErrDatabase
	}
}

func (p *postRepository) Delete(post *domain.Post) error {
	log.Print("[PostRepository]...Delete")
	result := p.DB.Delete(&post)

	switch result.Error {
	case nil:
		return nil
	case gorm.ErrRecordNotFound:
		return errors.ErrNotFound
	default:
		return errors.ErrDatabase
	}
}

func (p *postRepository) Migrate() error {
	log.Print("[PostRepository]...Migrate")
	result := p.DB.AutoMigrate(&domain.Post{})

	switch result.Error {
	case nil:
		return nil
	default:
		return errors.ErrDatabase
	}
}
