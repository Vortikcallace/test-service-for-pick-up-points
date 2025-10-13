package services

import (
	"errors"
	database "test-service-for-pick-up-points/internal/database"
	"test-service-for-pick-up-points/internal/models"
)

type OrderService struct {
	orderRepo   *database.OrderRepository
	userRepo    *database.UserRepository
	productRepo *database.ProductRepository
}

func NewOrderService(orderRepo *database.OrderRepository, userRepo *database.UserRepository, productRepo *database.ProductRepository) *OrderService {
	return &OrderService{
		orderRepo:   orderRepo,
		userRepo:    userRepo,
		productRepo: productRepo,
	}
}

func (s *OrderService) CreateOrder(req models.CreateOrderRequest) (*models.Order, error) {
	_, err := s.userRepo.GetByID(req.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	_, err = s.productRepo.GetByID(req.ProductID)
	if err != nil {
		return nil, errors.New("product not found")
	}

	order := &models.Order{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Readiness: false,
	}

	if err := s.orderRepo.Create(order); err != nil {
		return nil, err
	}

	return s.orderRepo.GetByID(order.ID)
}

func (s *OrderService) GetUserOrders(userID uint) ([]models.Order, error) {
	return s.orderRepo.GetByUserID(userID)
}

func (s *OrderService) UpdateOrderReadiness(id uint, readiness bool) error {
	return s.orderRepo.UpdateReadiness(id, readiness)
}
