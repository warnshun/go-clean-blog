package controllers

import (
	"net/http"

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

func (c PostController) GetAllPosts(ctx *gin.Context) {
	posts, err := c.service.GetAllPosts()
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
