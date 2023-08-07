package repository

import (
	"gin-socmed/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	Create(post *entity.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewPostRepository(db *gorm.DB) *postRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(post *entity.Post) error {
	err := r.db.Create(&post).Error

	return err
}
