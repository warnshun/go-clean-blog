package services

import (
	"go-clean-blog/lib"
	"go-clean-blog/models"
	"go-clean-blog/repository"

	"gorm.io/gorm"
)

// UserService service layer
type UserService struct {
	logger     lib.Logger
	repository repository.UserRepository
}

// NewUserService creates a new userservice
func NewUserService(logger lib.Logger, repository repository.UserRepository) UserService {
	return UserService{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s UserService) WithTrx(trxHandle *gorm.DB) UserService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// GetUser gets one user
func (s UserService) GetUser(id uint) (user models.User, err error) {
	return user, s.repository.First(&user, "id = ?", id).Error
}

// GetUser gets one user by username
func (s UserService) GetUserByUsername(username string) (user models.User, err error) {
	return user, s.repository.Preload("Password").First(&user, "username = ?", username).Error
}

// GetAllUsers get all the user
func (s UserService) GetAllUsers() (users []models.User, err error) {
	return users, s.repository.Find(&users).Error
}

// CreateUser call to create the user
func (s UserService) CreateUser(user *models.User) error {
	return s.repository.Create(&user).Error
}
