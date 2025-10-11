package models

import (
	"time"

	"gorm.io/gorm"
)

type Point struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Open_or_Close bool           `gorm:"default:false" json:"op_or_cl"`
	Changed       bool           `gorm:"default:false" json:"changed"`
	Address       string         `gorm:"type:varchar(255);not null" json:"address"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreatePointRequest struct {
	Address       string `json:"address" binding:"required"`
	Open_or_Close bool   `json:"open_or_close" binding:"required"`
}
