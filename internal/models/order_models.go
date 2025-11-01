package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null" json:"user_id"`
	ProductID uint           `gorm:"not null" json:"product_id"`
	Product   Product        `gorm:"foreignKey:ProductID" json:"product"`
	Readiness bool           `gorm:"default:false" json:"readiness"`
	Access    bool           `gorm:"default:false" json:"access"`
	Active    bool           `gorm:"default:true" json:"active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	PointID   uint           `gorm:"not null" json:"point_id"`
	Point     Point          `gorm:"foreignKey:PointID" json:"point"`
}

type CreateOrderRequest struct {
	UserID    uint `json:"user_id" binding:"required"`
	ProductID uint `json:"product_id" binding:"required"`
	PointID   uint `json:"point_id" binding:"required"`
}
