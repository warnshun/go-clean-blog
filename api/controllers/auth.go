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
	logger          lib.Logger
	service         services.Auth
	userService     services.User
	passwordService services.Password
}

// NewAuth creates new controller
func NewAuth(
	logger lib.Logger,
	service services.Auth,
	userService services.User,
	passwordService services.Password,
) Auth {
	return Auth{
		logger:          logger,
		service:         service,
		userService:     userService,
		passwordService: passwordService,
	}
}

// SignIn signs in user
func (c Auth) SignIn(ctx *gin.Context) {
	c.logger.Info("SignIn route called")
	// Currently not checking for username and password
	// Can add the logic later if necessary.
	user, _ := c.userService.GetOneUser(uint(1))
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

	ps := models.Password{
		UserID:   user.ID,
		Password: register.Password,
	}
	// create password
	if err := c.passwordService.
		WithTrx(trxHandle).
		CreatePassword(&ps); err != nil {
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
