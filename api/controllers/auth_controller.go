package controllers

import (
	"net/http"

	"go-clean-blog/api/apitool"

	"go-clean-blog/dtos"
	"go-clean-blog/lib"
	"go-clean-blog/models"
	"go-clean-blog/services"

	"github.com/gin-gonic/gin"
)

// AuthController struct
type AuthController struct {
	logger      lib.Logger
	service     services.AuthService
	userService services.UserService
}

// NewAuthController creates new controller
func NewAuthController(
	logger lib.Logger,
	service services.AuthService,
	userService services.UserService,
) AuthController {
	return AuthController{
		logger:      logger,
		service:     service,
		userService: userService,
	}
}

// Login signs in user
func (c AuthController) Login(ctx *gin.Context) {
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
func (c AuthController) Register(ctx *gin.Context) {
	var register dtos.UserRegister

	if err := ctx.ShouldBindJSON(&register); err != nil {
		c.logger.Error(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// trxHandle := ctx.MustGet(constants.DBTransaction).(*gorm.DB)
	trxHandle := apitool.GetTx(ctx)

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

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "register successfully",
	})
}
