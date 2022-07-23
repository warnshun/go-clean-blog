package services

import (
	"go-clean-blog/dtos"
	"go-clean-blog/lib"
	"go-clean-blog/models"
	"go-clean-blog/repository"

	"gorm.io/gorm"
)

type PostService struct {
	logger         lib.Logger
	repository     repository.PostRepository
	likeRepository repository.PostLikeRepository
}

func NewPostService(
	logger lib.Logger,
	repository repository.PostRepository,
	likeRepository repository.PostLikeRepository,
) PostService {
	return PostService{
		logger:         logger,
		repository:     repository,
		likeRepository: likeRepository,
	}
}

func (s PostService) WithTrx(trxHandle *gorm.DB) PostService {
	s.repository = s.repository.WithTrx(trxHandle)
	s.likeRepository = s.likeRepository.WithTrx(trxHandle)
	return s
}

func (s PostService) GetPost(id uint) (post models.Post, err error) {
	return post, s.repository.Preload("Photos").Preload("Author").First(&post, "id = ?", id).Error
}

func (s PostService) GetAllPosts(queryParam dtos.PostQuery) (posts []models.Post, err error) {
	query := s.repository.Preload("Photos").Preload("Author")
	if queryParam.UserID != nil {
		query = query.Where("user_id = ?", queryParam.UserID)
	}
	return posts, query.Find(&posts).Error
}

func (s PostService) GetAllPostsByUserID(userID uint) (posts []models.Post, err error) {
	return posts, s.repository.Preload("Photos").Preload("Author").Find(&posts, "user_id = ?", userID).Error
}

func (s PostService) GetAllLikedPostsByUserID(userID uint) (likedPost []models.PostLike, err error) {
	return likedPost, s.likeRepository.
		Preload("Post.Photos").
		Preload("Post.Author").
		Preload("Post").
		Find(&likedPost, "user_id = ?", userID).
		Error
}

func (s PostService) CreatePost(post *models.Post) error {
	return s.repository.Create(&post).Error
}

func (s PostService) GetPostLikeByPostIDandUserID(postID, userID uint) (postLike models.PostLike, err error) {
	return postLike,
		s.likeRepository.
			Where("post_id = ?", postID).
			Where("user_id = ?", userID).
			First(&postLike).
			Error
}

func (s PostService) CreatePostLike(postLike *models.PostLike) error {
	return s.likeRepository.Create(&postLike).Error
}

func (s PostService) DeletePostLike(id uint) error {
	return s.likeRepository.Delete(&models.PostLike{}, "id = ?", id).Error
}
