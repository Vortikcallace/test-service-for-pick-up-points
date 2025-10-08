package services

import "test-service-for-pick-up-points/internal/database"

type UserService struct {
	userRepo *database.UserRepository
}
