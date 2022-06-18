package models

import (
	"time"

	"gorm.io/gorm"
)

type PostPhoto struct {
	ID uint `gorm:"column:id;autoIncrement"`

	PostID uint   `gorm:"column:post_id"`
	Url    string `gorm:"column:url"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at"`

	Post Post `gorm:"foreignKey:PostID;references:ID"`
}

func (PostPhoto) TableName() string {
	return "post_photos"
}
