package apitool

import (
	"sync"

	"github.com/dipeshdulal/clean-gin/constants"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LazyTx struct {
	logger lib.Logger
	db     *gorm.DB
	tx     *gorm.DB
	once   sync.Once
}

func NewLazyTx(
	logger lib.Logger,
	db *gorm.DB,
) *LazyTx {
	lazyTx := &LazyTx{
		logger: logger,
		db:     db,
		once:   sync.Once{},
	}

	return lazyTx
}

func (lazyTx LazyTx) IsOpen() bool {
	return lazyTx.tx != nil
}

func (lazyTx *LazyTx) beginTx() {
	lazyTx.once.Do(func() {
		lazyTx.tx = lazyTx.db.Begin()
		lazyTx.logger.Info("beginning database transaction")

		defer func() {
			if r := recover(); r != nil {
				lazyTx.tx.Rollback()
			}
		}()
	})
}

func (lazyTx *LazyTx) getTx() *gorm.DB {
	if !lazyTx.IsOpen() {
		lazyTx.beginTx()
	}
	return lazyTx.tx
}

func SetTx(ctx *gin.Context, lazyTx *LazyTx) {
	ctx.Set(constants.DBTransaction, lazyTx)
}

func GetTx(ctx *gin.Context) *gorm.DB {
	lazyTx, exists := ctx.Get(constants.DBTransaction)
	if exists {
		return lazyTx.(*LazyTx).getTx()
	}
	return nil
}
