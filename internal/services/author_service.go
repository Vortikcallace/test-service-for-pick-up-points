package services

import (
	"errors"
	"test-service-for-pick-up-points/internal/database"
	"test-service-for-pick-up-points/internal/models"
	"test-service-for-pick-up-points/pkg/utils"
	"time"
)

type AuthorService struct {
	authorRepo *database.AuthorRepository
}

func NewAuthorService(authorRepo *database.AuthorRepository) *AuthorService {
	return &AuthorService{authorRepo: authorRepo}
}

func (s *AuthorService) CreateAuthor(req models.CreateAuthorRequest) (*models.AuthorResponse, error) {
	existingAuthor, _ := s.authorRepo.GetByEmail(req.Email)
	if existingAuthor != nil {
		return nil, errors.New("user with this email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	uniqueCode := uint(time.Now().Unix())

	author := &models.Author{
		Name:       req.Name,
		Email:      req.Email,
		Password:   hashedPassword,
		UniqueCode: uniqueCode,
	}

	if err := s.authorRepo.Create(author); err != nil {
		return nil, err
	}

	return &models.AuthorResponse{
		ID:         author.ID,
		Name:       author.Name,
		Email:      author.Email,
		UniqueCode: author.UniqueCode,
		CreatedAt:  author.CreatedAt,
	}, nil
}

func (s *AuthorService) GetAuthorProducts(id uint) (*models.Author, error) {
	author, err := s.authorRepo.GetAuthorProducts(id)
	if err != nil {
		return nil, errors.New("author not found")
	}
	return author, nil
}

func (s *AuthorService) GetAuthorProduct(id uint) (*models.Author, error) {
	author, err := s.authorRepo.GetAuthorProduct(id)
	if err != nil {
		return nil, errors.New("author not found")
	}
	return author, nil
}

func (s *AuthorService) GetAuthor(id uint) (*models.Author, error) {
	return s.authorRepo.GetByID(id)
}

func (s *AuthorService) UpdateAuthor(author *models.Author) error {
	return s.authorRepo.Update(author)
}

func (s *AuthorService) DeleteAuthor(author *models.Author) error {
	return s.authorRepo.Delete(author)
}
