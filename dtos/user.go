package dtos

import (
	"strconv"
	"time"
)

type User struct {
	ID          uint   `json:"id"`
	Username    string `json:"username"`
	Nickname    string `json:"nickname"`
	CreatedTime string `json:"created_time"`
	UpdatedTime string `json:"updated_time"`
}

func (m *User) CreatedAt(time time.Time) {
	m.CreatedTime = strconv.FormatInt(time.UnixMilli(), 10)
}

func (m *User) UpdatedAt(time time.Time) {
	m.UpdatedTime = strconv.FormatInt(time.UnixMilli(), 10)
}
