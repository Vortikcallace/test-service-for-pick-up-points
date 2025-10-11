package database

import (
	"test-service-for-pick-up-points/internal/models"

	"gorm.io/gorm"
)

type PointRepository struct {
	db *gorm.DB
}

func NewPointRepository(db *gorm.DB) *PointRepository {
	return &PointRepository{db: db}
}

func (r *PointRepository) Create(point *models.Point) error {
	return r.db.Create(point).Error
}

func (r *PointRepository) GetByID(id uint) (*models.Point, error) {
	var point models.Point
	err := r.db.First(&point, id).Error
	return &point, err
}

func (r *PointRepository) GetAll() ([]models.Point, error) {
	var points []models.Point
	err := r.db.Find(&points).Error
	return points, err
}
