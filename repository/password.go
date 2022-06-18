package repository

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"gorm.io/gorm"
)

type Password struct {
	lib.Database
	logger lib.Logger
}

func NewPassword(db lib.Database, logger lib.Logger) Password {
	return Password{
		Database: db,
		logger:   logger,
	}
}

func (r Password) WithTrx(trxHandle *gorm.DB) Password {
	if trxHandle == nil {
		r.logger.Error("Transaction Database not found in gin context. ")
		return r
	}
	r.Database.DB = trxHandle
	return r
}
