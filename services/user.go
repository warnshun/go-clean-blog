package services

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"gorm.io/gorm"
)

// User service layer
type User struct {
	logger     lib.Logger
	repository repository.UserRepository
}

// NewUser creates a new userservice
func NewUser(logger lib.Logger, repository repository.UserRepository) User {
	return User{
		logger:     logger,
		repository: repository,
	}
}

// WithTrx delegates transaction to repository database
func (s User) WithTrx(trxHandle *gorm.DB) User {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

// GetUser gets one user by username
func (s User) GetUserByUsername(username string) (user models.User, err error) {
	return user, s.repository.Preload("Password").Find(&user, "username = ?", username).Error
}

// GetAllUser get all the user
func (s User) GetAllUser() (users []models.User, err error) {
	return users, s.repository.Find(&users).Error
}

// CreateUser call to create the user
func (s User) CreateUser(user *models.User) error {
	return s.repository.Create(&user).Error
}

// UpdateUser updates the user
func (s User) UpdateUser(user models.User) error {
	return s.repository.Save(&user).Error
}

// DeleteUser deletes the user
func (s User) DeleteUser(id uint) error {
	return s.repository.Delete(&models.User{}, id).Error
}
