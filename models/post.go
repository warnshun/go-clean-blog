package models

import (
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID uint `gorm:"column:id;autoIncrement"`

	UserID  uint   `gorm:"column:user_id"`
	Content string `gorm:"column:content"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Photos []PostPhoto `gorm:"foreignKey:PostID;references:ID"`
	Author User        `gorm:"foreignKey:ID;references:UserID"`
}

func (Post) TableName() string {
	return "posts"
}
