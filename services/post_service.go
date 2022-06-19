package services

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/repository"
	"gorm.io/gorm"
)

type PostService struct {
	logger     lib.Logger
	repository repository.PostRepository
}

func NewPostService(logger lib.Logger, repository repository.PostRepository) PostService {
	return PostService{
		logger:     logger,
		repository: repository,
	}
}

func (s PostService) WithTrx(trxHandle *gorm.DB) PostService {
	s.repository = s.repository.WithTrx(trxHandle)
	return s
}

func (s PostService) GetPost(id uint) (post models.Post, err error) {
	return post, s.repository.Preload("Photos").Preload("Author").First(&post, "id = ?", id).Error
}

func (s PostService) GetAllPosts() (posts []models.Post, err error) {
	return posts, s.repository.Preload("Photos").Preload("Author").Find(&posts).Error
}

func (s PostService) CreatePost(post *models.Post) error {
	return s.repository.Create(&post).Error
}
