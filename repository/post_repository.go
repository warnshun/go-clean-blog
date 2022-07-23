package repository

import (
	"go-clean-blog/lib"

	"gorm.io/gorm"
)

type PostRepository struct {
	lib.Database
	logger lib.Logger
}

func NewPostRepository(db lib.Database, logger lib.Logger) PostRepository {
	return PostRepository{
		Database: db,
		logger:   logger,
	}
}

func (r PostRepository) WithTrx(trxHandle *gorm.DB) PostRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
