package services

import (
	"test-service-for-pick-up-points/internal/database"
	"test-service-for-pick-up-points/internal/models"
)

type ProductService struct {
	productRepo *database.ProductRepository
}

func NewProductService(productRepo *database.ProductRepository) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (s *ProductService) CreateProduct(req models.CreateProductRequest) (*models.Product, error) {
	product := &models.Product{
		Author:           req.Author,
		Name:             req.Name,
		ShortDescription: req.ShortDescription,
	}

	if err := s.productRepo.Create(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.productRepo.GetAll()
}

func (s *ProductService) GetProduct(id uint) (*models.Product, error) {
	return s.productRepo.GetByID(id)
}
