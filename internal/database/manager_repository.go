package database

import (
	"test-service-for-pick-up-points/internal/models"

	"gorm.io/gorm"
)

type ManagerRepository struct {
	db *gorm.DB
}

func NewManagerRepository(db *gorm.DB) *ManagerRepository {
	return &ManagerRepository{db: db}
}

func (r *ManagerRepository) Create(manager *models.Manager) error {
	return r.db.Create(manager).Error
}

func (r *ManagerRepository) GetByEmail(email string) (*models.Manager, error) {
	var manager models.Manager
	err := r.db.Where("email = ?", email).First(&manager).Error
	return &manager, err
}

func (r *ManagerRepository) GetByID(id uint) (*models.Manager, error) {
	var manager models.Manager
	err := r.db.Preload("Point").First(&manager, id).Error
	return &manager, err
}

func (r *ManagerRepository) GetManagerOrders(id uint) (*models.Manager, error) {
	var manager models.Manager
	err := r.db.Preload("Orders.Product").First(&manager, id).Error
	return &manager, err
}

func (r *ManagerRepository) GetAll() ([]models.Manager, error) {
	var managers []models.Manager
	err := r.db.Find(&managers).Error
	return managers, err
}

func (r *ManagerRepository) Update(manager *models.Manager) error {
	return r.db.Save(manager).Error
}

func (r *ManagerRepository) Delete(manager *models.Manager) error {
	return r.db.Delete(manager).Error
}
