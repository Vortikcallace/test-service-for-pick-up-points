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
}

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

type Product struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	Author           string         `gorm:"size:255" json:"author"`
	Name             string         `gorm:"size:20" json:"name"`
	ShortDescription string         `gorm:"size:20" json:"shortdescr"`
	CreatedAt        time.Time      `json:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"-"`
}
