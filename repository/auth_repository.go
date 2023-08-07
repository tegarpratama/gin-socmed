package repository

import (
	"gin-socmed/entity"

	"gorm.io/gorm"
)

type AuthRepository interface {
	Register(req *entity.User) error
	GetUserByEmail(email string) (*entity.User, error)
	EmailExists(email string) bool
}

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) *authRepository {
	return &authRepository{
		db: db,
	}
}

func (r *authRepository) Register(user *entity.User) error {
	err := r.db.Create(&user).Error

	return err
}

func (r *authRepository) GetUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, "email = ?", email).Error

	return &user, err
}

func (r *authRepository) EmailExists(email string) bool {
	var user entity.User
	err := r.db.First(&user, "email = ?", email).Error

	return err == nil
}
