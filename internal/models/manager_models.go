package models

import (
	"time"

	"gorm.io/gorm"
)

type Manager struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"size:20" json:"name"`
	Email     string         `gorm:"type:varchar(100);uniqueIndex" json:"email"`
	Password  string         `gorm:"size:255" json:"-"`
	Point     Point          `gorm:"foreignKey:PointID" json:"point"`
	PointID   uint           `gorm:"not null" json:"point_id"`
	Orders    []Order        `gorm:"foreignKey:UserID" json:"orders"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type UpdateManagerPasswordRequest struct {
	OldPassword string `json:"old_pw" validate:"required"`
	NewPaaword  string `json:"new_pw" validate:"required,min=6"`
}

type ManagerResponse struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email" `
	Orders    []Order   `json:"orders"`
	CreatedAt time.Time `json:"created_at"`
	Point     Point     `json:"point"`
}

type CreateManagerRequests struct {
	Name     string  `json:"name" binding:"required,min=2,max=20"`
	Email    string  `json:"email" binding:"required,email"`
	Password string  `json:"password" binding:"required,min=6"`
	Orders   []Order `json:"orders" binding:"required"`
	Point    Point   `json:"point" binding:"required"`
}
