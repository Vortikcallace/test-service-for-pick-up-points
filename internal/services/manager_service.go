package services

import (
	"errors"
	"test-service-for-pick-up-points/internal/database"
	"test-service-for-pick-up-points/internal/models"
	"test-service-for-pick-up-points/pkg/utils"
)

type ManagerService struct {
	managerRepo *database.ManagerRepository
}

func NewManagerService(managerRepo *database.ManagerRepository) *ManagerService {
	return &ManagerService{managerRepo: managerRepo}
}

func (s *ManagerService) CreateManager(req models.CreateManagerRequests) (*models.ManagerResponse, error) {
	existingManager, _ := s.managerRepo.GetByEmail(req.Email)
	if existingManager != nil {
		return nil, errors.New("manager with this email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	manager := &models.Manager{
		Name:     req.Name,
		Password: hashedPassword,
		Email:    req.Email,
		Orders:   req.Orders,
		Point:    req.Point,
	}

	if err := s.managerRepo.Create(manager); err != nil {
		return nil, err
	}

	return &models.ManagerResponse{
		ID:        manager.ID,
		Name:      manager.Name,
		Email:     req.Email,
		CreatedAt: manager.CreatedAt,
		Orders:    manager.Orders,
		Point:     manager.Point,
	}, nil
}

func (s *ManagerService) GetManagerOrders(id uint) (*models.Manager, error) {
	manager, err := s.managerRepo.GetManagerOrders(id)
	if err != nil {
		return nil, errors.New("manager not found")
	}
	return manager, nil
}

func (s *ManagerService) GetManagerPoint(id uint) (*models.Manager, error) {
	manager, err := s.managerRepo.GetManagerPoint(id)
	if err != nil {
		return nil, errors.New("manager not found")
	}
	return manager, nil
}

func (s *ManagerService) GetManager(id uint) (*models.Manager, error) {
	return s.managerRepo.GetByID(id)
}

func (s *ManagerService) UpdateManager(manager *models.Manager) error {
	return s.managerRepo.Update(manager)
}

func (s *ManagerService) DeleteManager(manager *models.Manager) error {
	return s.managerRepo.Delete(manager)
}
