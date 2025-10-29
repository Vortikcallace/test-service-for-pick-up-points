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

/*
I will add some functions comig soon
*/

func (r *AuthorRepository) Update(author *models.Author) error {
	return r.db.Save(author).Error
}
