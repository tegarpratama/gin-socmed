package repository

import (
	"fmt"
	"gin-socmed/dto"
	"gin-socmed/entity"

	"gorm.io/gorm"
)

type UserRepository interface {
	TotalUsers(params *dto.FilterParams) (int64, error)
	GetUsers(params *dto.FilterParams) (*[]dto.UserResponse, error)
	Detail(id int) (*dto.UserResponse, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) TotalUsers(params *dto.FilterParams) (int64, error) {
	var count int64
	query := r.db.Model(&entity.User{})

	if params.Search != "" {
		query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", params.Search)).Or("email LIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	err := query.Count(&count)
	if err != nil {
		return count, err.Error
	}

	return count, nil
}

func (r *userRepository) GetUsers(params *dto.FilterParams) (*[]dto.UserResponse, error) {
	var userResponse []dto.UserResponse
	query := r.db.Model(&entity.User{}).Select("id, name, email, gender, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') as updated_at")

	if params.Search != "" {
		query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", params.Search)).Or("email LIKE ?", fmt.Sprintf("%%%s%%", params.Search))
	}

	err := query.Preload("Post", func(db *gorm.DB) *gorm.DB {
		db = db.Select("id, user_id, tweet, picture_url, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') as updated_at").Order("id desc")
		return db
	}).Offset(params.Offset).Limit(params.Limit).Find(&userResponse).Error

	return &userResponse, err
}

func (r *userRepository) Detail(id int) (*dto.UserResponse, error) {
	var usersResponse dto.UserResponse
	err := r.db.Model(&entity.User{}).Select("id, name, email, gender, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') as updated_at").Preload("Post", func(db *gorm.DB) *gorm.DB {
		db = db.Select("id, user_id, tweet, picture_url, DATE_FORMAT(created_at, '%Y-%m-%d %H:%i:%s') as created_at, DATE_FORMAT(updated_at, '%Y-%m-%d %H:%i:%s') as updated_at").Order("id desc")
		return db
	}).First(&usersResponse, id).Error

	return &usersResponse, err
}
