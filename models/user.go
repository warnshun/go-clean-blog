package models

import (
	"time"

	"gorm.io/gorm"
)

// User model
type User struct {
	ID uint `gorm:"column:id;autoIncrement"`

	Username string `gorm:"column:username"`
	Nickname string `gorm:"column:nickname"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

// TableName gives table name of model
func (User) TableName() string {
	return "users"
}
