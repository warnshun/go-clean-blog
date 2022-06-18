package controllers

import (
	"net/http"

	"github.com/dipeshdulal/clean-gin/constants"
	"github.com/dipeshdulal/clean-gin/dtos"
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/models"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Auth struct
type Auth struct {
	logger      lib.Logger
	service     services.Auth
	userService services.User
}

// NewAuth creates new controller
func NewAuth(
	logger lib.Logger,
	service services.Auth,
	userService services.User,
) Auth {
	return Auth{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// Login signs in user
func (c Auth) Login(ctx *gin.Context) {
	// Currently not checking for username and password
	// Can add the logic later if necessary.
	var login dtos.UserLogin
	if err := ctx.ShouldBindJSON(&login); err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := c.userService.GetUserByUsername(login.Username)
	if err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	if login.Password != user.Password.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "password is incorrect",
		})
		return
	}

	token := c.service.CreateToken(user)
	ctx.JSON(200, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}

// Register registers user
func (c Auth) Register(ctx *gin.Context) {
	var register dtos.UserRegister

	if err := ctx.ShouldBindJSON(&register); err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	trxHandle := ctx.MustGet(constants.DBTransaction).(*gorm.DB)

	user := models.User{
		Username: register.Username,
		Password: models.Password{
			Password: register.Password,
		},
	}

	// create user
	if err := c.userService.
		WithTrx(trxHandle).
		CreateUser(&user); err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "register successfully",
	})
}
