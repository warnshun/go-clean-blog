package repository

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"gorm.io/gorm"
)

type PostLikeRepository struct {
	lib.Database
	logger lib.Logger
}

func NewPostLikeRepository(
	db lib.Database,
	logger lib.Logger,
) PostLikeRepository {
	return PostLikeRepository{
		Database: db,
		logger:   logger,
	}
}

func (r PostLikeRepository) WithTrx(trxHandle *gorm.DB) PostLikeRepository {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
