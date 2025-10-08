package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Author           string         `gorm:"size:255" json:"author"`
	Name             string         `gorm:"size:20" json:"name"`
	ShortDescription string         `gorm:"size:20" json:"shortdescr"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
