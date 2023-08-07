package repository

import (
	"errors"
	"fmt"
	"gin-socmed/dto"
	"gin-socmed/entity"

	"gorm.io/gorm"
)

type PostRepository interface {
	TotalData(params *dto.FilterParams) (int64, error)
	FindAll(params *dto.FilterParams) (*[]dto.PostsResponse, error)
	Detail(id *int) (dto.PostsResponse, error)
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

func (r *postRepository) TotalData(params *dto.FilterParams) (int64, error) {
	var count int64
	query := r.db.Model(&entity.Post{})

	if params.Search != "" {
		query.Where("tweet LIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	err := query.Count(&count)
	if err != nil {
		return count, err.Error
	}

	return count, nil
}

func (r *postRepository) FindAll(params *dto.FilterParams) (*[]dto.PostsResponse, error) {
	var postsResponse []dto.PostsResponse
	query := r.db.Model(&entity.Post{}).Select("id, user_id,tweet, picture_url, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') as updated_at")

	if params.Search != "" {
		query.Where("tweet LIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	err := query.Preload("User").Offset(params.Offset).Limit(params.Limit).Find(&postsResponse).Error

	return &postsResponse, err
}

func (r *postRepository) Detail(id *int) (dto.PostsResponse, error) {
	var postsResponse dto.PostsResponse
	err := r.db.Model(&entity.Post{}).Select("id, user_id,tweet, picture_url, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') as updated_at").Preload("User").First(&postsResponse, id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("tweet not found")
	}

	return postsResponse, err
}

func (r *postRepository) Create(post *entity.Post) error {
	err := r.db.Create(&post).Error

	return err
}
