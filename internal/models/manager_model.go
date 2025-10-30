package models

import (
	"time"

	"gorm.io/gorm"
)

type Manager struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:20" json:"name"`
	Password  string         `gorm:"size:255" json:"-"`
	Point     Point          `gorm:"" json:""`
	PointID   uint           `gorm:"" json:""`
	Orders    []Order        `gorm:"foreignKey:UserID" json:"orders"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type ManagerResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Orders    []Order   `json:"orders"`
	CreatedAt time.Time `json:"created_at"`
	Point     Point     `json:"point"`
}

type CreateManagerRequests struct {
	Name     string `json:"name" binding:"required,min=2,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}
