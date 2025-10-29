package services

import (
	"errors"
	"test-service-for-pick-up-points/internal/database"
	"test-service-for-pick-up-points/internal/models"
	"test-service-for-pick-up-points/pkg/utils"
	"time"
)

type UserService struct {
	userRepo *database.UserRepository
}

func NewUserService(userRepo *database.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(req models.CreateUserRequest) (*models.UserResponse, error) {
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	uniqueCode := uint(time.Now().Unix())

	user := &models.User{
		Name:       req.Name,
		Email:      req.Email,
		Password:   hashedPassword,
		Age:        req.Age,
		UniqueCode: uniqueCode,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return &models.UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Age:        user.Age,
		UniqueCode: user.UniqueCode,
		CreatedAt:  user.CreatedAt,
	}, nil
}

func (s *UserService) GetUserWithOrders(id uint) (*models.User, error) {
	user, err := s.userRepo.GetWithOrders(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) GetUser(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}

func (s *UserService) UpdateUser(user *models.User) error {
	return s.userRepo.Update(user)
}

//
//
//
//
//Unusable function
/*
func (s *UserService) Authenticate(email, password string) (*models.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
*/
