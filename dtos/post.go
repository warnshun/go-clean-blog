package dtos

import (
	"strconv"
	"time"

	"go-clean-blog/models"
)

type Post struct {
	ID          uint     `json:"id"`
	PhotoUrls   []string `json:"photo_urls"`
	Content     string   `json:"content"`
	Author      User     `json:"author"`
	CreatedTime string   `json:"created_time"`
}

func (m *Post) CreatedAt(time time.Time) {
	m.CreatedTime = strconv.FormatInt(time.UnixMilli(), 10)
}

func (m *Post) Photos(photos []models.PostPhoto) {
	urls := make([]string, 0, len(photos))
	for _, photo := range photos {
		urls = append(urls, photo.Url)
	}
	m.PhotoUrls = urls
}
