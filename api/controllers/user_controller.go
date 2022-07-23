package controllers

import (
	"net/http"
	"strconv"

	"go-clean-blog/dtos"
	"go-clean-blog/lib"
	"go-clean-blog/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// UserController data type
type UserController struct {
	logger  lib.Logger
	service services.UserService
}

// NewUserController creates new user controller
func NewUserController(userService services.UserService, logger lib.Logger) UserController {
	return UserController{
		service: userService,
		logger:  logger,
	}
}

// GetOneUser gets one user
func (u UserController) GetOneUser(c *gin.Context) {
	paramID := c.Param("id")

	id, err := strconv.Atoi(paramID)
	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user, err := u.service.GetUser(uint(id))

	var dto dtos.User
	copier.Copy(&dto, user)

	if err != nil {
		u.logger.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": dto,
	})

}

// GetAllUsers gets the user
func (u UserController) GetAllUsers(c *gin.Context) {
	users, err := u.service.GetAllUsers()
	if err != nil {
		u.logger.Error(err)
	}

	var dtos []dtos.User
	copier.Copy(&dtos, users)
	c.JSON(200, gin.H{"data": dtos})
}
