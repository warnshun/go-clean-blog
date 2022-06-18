package dtos

type PostAdd struct {
	PhotoUrls []string `json:"photo_urls"`
	Content   string   `json:"content"`
}
