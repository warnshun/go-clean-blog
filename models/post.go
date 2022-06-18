package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID uint `gorm:"column:id;autoIncrement"`

	Content string `gorm:"column:content"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`
}

func (Post) TableName() string {
	return "posts"
}
