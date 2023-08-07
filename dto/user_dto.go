package dto

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	CreatedAt string `json:"created_at"`
	Post      []Post `gorm:"foreignKey:UserID" json:"tweets"`
}

type Post struct {
	UserID    int    `json:"-"`
	ID        int    `json:"id"`
	Tweet     string `json:"tweet"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
