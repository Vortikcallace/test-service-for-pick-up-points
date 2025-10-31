package database

import (
	"test-service-for-pick-up-points/internal/models"

	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}

func (r *AuthorRepository) Create(author *models.Author) error {
	return r.db.Create(author).Error
}

func (r *AuthorRepository) GetByEmail(email string) (*models.Author, error) {
	var author models.Author
	err := r.db.Where("email = ?", email).First(&author).Error
	return &author, err
}

func (r *AuthorRepository) GetByID(id uint) (*models.Author, error) {
	var author models.Author
	err := r.db.Preload("Products").First(&author, id).Error
	return &author, err
}

func (r *AuthorRepository) GetAll() ([]models.Author, error) {
	var authors []models.Author
	err := r.db.Find(authors).Error
	return authors, err
}

func (r *AuthorRepository) Update(author *models.Author) error {
	return r.db.Save(author).Error
}

func (r *AuthorRepository) Delete(author *models.Author) error {
	return r.db.Delete(author).Error
}
