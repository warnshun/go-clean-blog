package models

import (
	"time"

	"gorm.io/gorm"
)

type Password struct {
	UserID   uint   `gorm:"column:user_id"`
	Password string `gorm:"column:password"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Password) TableName() string {
	return "password"
}
