package services

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"gorm.io/gorm"
)

type Password struct {
	logger     lib.Logger
	repository repository.Password
}

func NewPassword(logger lib.Logger, repository repository.Password) Password {
	return Password{
		logger:     logger,
		repository: repository,
	}
}

func (s Password) WithTrx(trxHandle *gorm.DB) Password {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s Password) CreatePassword(password *models.Password) error {
	return s.repository.Create(&password).Error
}
