package repositories

import (
	"go-todo/models"

	"gorm.io/gorm"
)

type PostRepository interface {
	GetAllPosts() ([]models.Post, error)
	GetPostByID(ID int) (*models.Post, error)
	CreatePost(post models.Post) (models.Post, error)
	UpdatePost(post *models.Post) error
	DeletePost(ID int) error
}

func RepositoryPost(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) CreatePost(post models.Post) (models.Post, error) {
	err := r.db.Create(&post).Error

	return post, err
}

func (r *repository) GetAllPosts() ([]models.Post, error) {
	var post []models.Post
	err := r.db.Find(&post).Error
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (r *repository) GetPostByID(ID int) (*models.Post, error) {
	var post models.Post
	err := r.db.First(&post, ID).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *repository) UpdatePost(post *models.Post) error {
	return r.db.Save(post).Error
}

func (r *repository) DeletePost(ID int) error {
	var post models.Post
	err := r.db.First(&post, ID).Error
	if err != nil {
		return err
	}

	err = r.db.Delete(&post).Error
	if err != nil {
		return err
	}

	return nil
}
