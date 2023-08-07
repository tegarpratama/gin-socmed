package entity

import (
	"time"
)

type Post struct {
	ID         int
	UserID     int
	Tweet      string
	PictureUrl *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
