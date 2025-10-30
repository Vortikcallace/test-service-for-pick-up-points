package models

import (
	"time"

	"gorm.io/gorm"
)

type Author struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"size:20" json:"name"`
	Email      string         `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password   string         `gorm:"size:255" json:"-"`
	UniqueCode uint           `gorm:"primaryKey" json:"ucode"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Products   []Product      `gotm:"" json:""`
}

type UpdateAuthorPasswordRequest struct {
	OldPassword string `json:"old_pw" validate:"required"`
	NewPaaword  string `json:"new_pw" validate:"required,min=6"`
}

type AuthorResponse struct {
	ID         uint      `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	UniqueCode uint      `json:"ucode"`
	CreatedAt  time.Time `json:"created_at"`
}

type CreateAuthorRequest struct {
	Name     string `json:"name" binding:"required,min=2,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
