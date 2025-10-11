package models

import (
	"time"

	"gorm.io/gorm"
)

type Point struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	ProductID     uint           `gorm:"not null" json:"product_id"`
	Open_or_Close bool           `gorm:"default:false" json:"readiness"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreatePointRequest struct {
	Address       string `json:"address" binding:"required"`
	Open_or_Close bool   `json:"open_or_close" binding:"required"`
}
