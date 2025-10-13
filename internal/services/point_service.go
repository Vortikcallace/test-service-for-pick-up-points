package services

import (
	"test-service-for-pick-up-points/internal/database"
	"test-service-for-pick-up-points/internal/models"
)

type PointService struct {
	pointRepo *database.PointRepository
}

func NewPointService(pointRepo *database.PointRepository) *PointService {
	return &PointService{pointRepo: pointRepo}
}

func (s *PointService) CreatePoint(req models.CreatePointRequest) (*models.Point, error) {
	point := &models.Point{
		Address:       req.Address,
		Changed:       req.Changed,
		Open_or_Close: req.Open_or_Close,
	}

	if err := s.pointRepo.Create(point); err != nil {
		return nil, err
	}

	return point, nil
}

func (s *PointService) GetAllPoints() ([]models.Point, error) {
	return s.pointRepo.GetAll()
}

func (s *PointService) GetPoint(id uint) (*models.Point, error) {
	return s.pointRepo.GetByID(id)
}
