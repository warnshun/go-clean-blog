package controllers

import (
	"github.com/dipeshdulal/clean-gin/lib"
	"github.com/dipeshdulal/clean-gin/services"
	"github.com/gin-gonic/gin"
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

// SignIn signs in user
func (jwt Auth) SignIn(c *gin.Context) {
	jwt.logger.Info("SignIn route called")
	// Currently not checking for username and password
	// Can add the logic later if necessary.
	user, _ := jwt.userService.GetOneUser(uint(1))
	token := jwt.service.CreateToken(user)
	c.JSON(200, gin.H{
		"message": "logged in successfully",
		"token":   token,
	})
}

// Register registers user
func (jwt Auth) Register(c *gin.Context) {
	jwt.logger.Info("Register route called")
	c.JSON(200, gin.H{
		"message": "register route222",
	})
}
