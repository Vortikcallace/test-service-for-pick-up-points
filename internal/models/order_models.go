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
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
