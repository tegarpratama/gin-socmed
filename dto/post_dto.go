package dto

import (
	"mime/multipart"
)

type PostsResponse struct {
	ID         int    `json:"id"`
	UserID     int    `json:"-"`
	User       User   `gorm:"foreignKey:UserID" json:"user"`
	Tweet      string `json:"tweet"`
	PictureUrl string `json:"picture_url"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type PostRequest struct {
	UserID  int                   `form:"user_id"`
	Tweet   string                `form:"tweet" binding:"required"`
	Picture *multipart.FileHeader `form:"picture"`
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
