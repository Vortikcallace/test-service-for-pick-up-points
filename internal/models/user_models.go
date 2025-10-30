package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"size:20" json:"name"`
	Email      string         `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password   string         `gorm:"size:255" json:"-"`
	Age        int            `gorm:"default:18" json:"age"`
	UniqueCode uint           `gorm:"primaryKey" json:"ucode"`
	Orders     []Order        `gorm:"foreignKey:UserID" json:"orders"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Point      Point
	AllPoints  []Point
}

type UpdatePasswordRequest struct {
	OldPassword string `json:"old_pw" validate:"required"`
	NewPaaword  string `json:"new_pw" validate:"required,min=6"`
}

type UserResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Age        int       `json:"age"`
	UniqueCode uint      `json:"ucode"`
	Orders     []Order   `json:"orders"`
	CreatedAt  time.Time `json:"created_at"`
	Point      Point     `json:"point"`
	AllPoints  []Point   `json:"allpoints"`
}

type CreateUserRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"min=0,max=120"`
}
