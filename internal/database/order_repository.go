package database

import (
	"test-service-for-pick-up-points/internal/models"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *OrderRepository) GetByID(id uint) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Product").First(&order, id).Error
	return &order, err
}

func (r *OrderRepository) GetByUserID(userID uint) ([]models.Order, error) {
	var orders []models.Order
	err := r.db.Preload("Product").Where("user_id = ?", userID).Find(&orders).Error
	return orders, err
}

func (r *OrderRepository) UpdateReadiness(id uint, readiness bool) error {
	return r.db.Model(&models.Order{}).Where("id = ?", id).Update("readiness", readiness).Error
}
