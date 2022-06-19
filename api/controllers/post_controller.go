package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/dipeshdulal/clean-gin/constants"
	"github.com/dipeshdulal/clean-gin/dtos"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
)

type PostController struct {
	logger  lib.Logger
	service services.PostService
}

func NewPostController(
	logger lib.Logger,
	postService services.PostService,
) PostController {
	return PostController{
		logger:  logger,
		service: postService,
	}
}

func (c PostController) GetPost(ctx *gin.Context) {
	paramID := ctx.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	post, err := c.service.GetPost(uint(id))
	if err != nil {
		c.logger.Error(err)
	}

	var dto dtos.Post
	copier.Copy(&dto, post)

	ctx.JSON(200, gin.H{"data": dto})
}

func (c PostController) GetAllPosts(ctx *gin.Context) {
	var queryParam dtos.PostQuery

	if err := ctx.ShouldBindQuery(&queryParam); err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	posts, err := c.service.GetAllPosts(queryParam)
	if err != nil {
		c.logger.Error(err)
	}

	var dtos []dtos.Post

	copier.Copy(&dtos, posts)

	ctx.JSON(200, gin.H{"data": dtos})
}

func (c PostController) AddPost(ctx *gin.Context) {
	var postAdd dtos.PostAdd

	if err := ctx.ShouldBindJSON(&postAdd); err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	photos := make([]models.PostPhoto, 0, len(postAdd.PhotoUrls))
	for _, photo := range postAdd.PhotoUrls {
		photos = append(photos, models.PostPhoto{
			Url: photo,
		})
	}

	token := ctx.MustGet(constants.JWTToken).(services.JWTToken)

	post := models.Post{
		UserID:  token.ID,
		Photos:  photos,
		Content: postAdd.Content,
	}

	trxHandle := ctx.MustGet(constants.DBTransaction).(*gorm.DB)

	if err := c.service.WithTrx(trxHandle).
		CreatePost(&post); err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "post created successfully",
	})
}

func (c PostController) SwitchLikePost(ctx *gin.Context) {
	var input dtos.PostLike
	if err := ctx.ShouldBindJSON(&input); err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	token := ctx.MustGet(constants.JWTToken).(services.JWTToken)

	trxHandle := ctx.MustGet(constants.DBTransaction).(*gorm.DB)
	svTx := c.service.WithTrx(trxHandle)

	postLike, err := svTx.GetPostLikeByPostIDandUserID(input.PostID, token.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		postLike.PostID = input.PostID
		postLike.UserID = token.ID
		if err := svTx.CreatePostLike(&postLike); err != nil {
			c.logger.Error(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"message": "post liked successfully",
		})

		return
	}

	if err := svTx.DeletePostLike(postLike.ID); err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "post unliked successfully",
	})
}

func (c PostController) GetAllLikedPosts(ctx *gin.Context) {
	token := ctx.MustGet(constants.JWTToken).(services.JWTToken)

	posts, err := c.service.GetAllLikedPostsByUserID(token.ID)
	if err != nil {
		c.logger.Error(err)
	}

	postDtos := make([]dtos.Post, 0, len(posts))
	for _, post := range posts {
		var dto dtos.Post
		copier.Copy(&dto, post.Post)
		postDtos = append(postDtos, dto)
	}

	ctx.JSON(http.StatusOK, gin.H{"data": postDtos})
}
