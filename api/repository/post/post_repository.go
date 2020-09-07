package repository

import (
	"fmt"
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
	result := p.DB.Create(&post)

	if result.Error != nil {
		err := result.Error
		msg := "error saving the post"
		err = errors.InternalError.Wrapf(err, msg)
		return nil, err
	}
	return post, nil
}

func (p *postRepository) FindAll() ([]domain.Post, error) {
	log.Print("[PostRepository]...FindAll")
	var posts []domain.Post
	result := p.DB.Find(&posts)

	if result.Error != nil {
		err := result.Error
		msg := "error getting the posts"
		switch err {
		case gorm.ErrRecordNotFound:
			err = errors.NotFound.Wrapf(err, msg)
		default:
			err = errors.InternalError.Wrapf(err, msg)
		}
		return nil, err
	}
	return posts, nil
}

func (p *postRepository) FindByID(id int64) (*domain.Post, error) {
	log.Print("[PostRepository]...FindAll")
	var post domain.Post
	result := p.DB.Where("id = ?", id).First(&post)

	if result.Error != nil {
		err := result.Error
		msg := fmt.Sprintf("error getting the post with id %d", id)
		switch err {
		case gorm.ErrRecordNotFound:
			err = errors.NotFound.Wrapf(err, msg)
		default:
			err = errors.InternalError.Wrapf(err, msg)
		}
		return nil, err
	}
	return &post, nil
}

func (p *postRepository) Delete(post *domain.Post) error {
	log.Print("[PostRepository]...Delete")
	result := p.DB.Delete(&post)

	if result.Error != nil {
		err := result.Error
		msg := "error deleting the post"
		err = errors.InternalError.Wrapf(err, msg)
		return err
	}
	return nil
}

func (p *postRepository) Migrate() error {
	log.Print("[PostRepository]...Migrate")
	result := p.DB.AutoMigrate(&domain.Post{})

	switch result.Error {
	case nil:
		return nil
	default:
		return result.Error
	}
}
