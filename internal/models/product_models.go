package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Author           Author         `gorm:"foreignKey:AuthorID" json:"author"`
	AuthorID         uint           `gorm:"not null" json:"author_id"`
	Name             string         `gorm:"size:20" json:"name"`
	ShortDescription string         `gorm:"size:20" json:"shortdescr"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}

type CreateProductRequest struct {
	Author           Author `json:"author" binding:"required"`
	Name             string `json:"name" binding:"required,max=20"`
	ShortDescription string `json:"shortdescr" binding:"required,max=20"`
}
