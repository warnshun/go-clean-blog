package models

import "time"

type PostLike struct {
	ID uint `gorm:"column:id;autoIncrement"`

	PostID uint `gorm:"column:post_id"`
	UserID uint `gorm:"column:user_id"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:milli"`

	Post Post `gorm:"foreignkey:ID;references:PostID"`
}

func (PostLike) TableName() string {
	return "post_likes"
}
