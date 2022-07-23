package routes

import (
	"go-clean-blog/api/controllers"
	"go-clean-blog/lib"
)

// AuthRoute struct
type AuthRoute struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	authController controllers.AuthController
}

// Setup user routes
func (s AuthRoute) Setup() {
	s.logger.Info("Setting up routes")
	auth := s.handler.Gin.Group("/auth")
	{
		auth.POST("/login", s.authController.Login)
		auth.POST("/register", s.authController.Register)
	}
}

// NewAuthRoute creates new user controller
func NewAuthRoute(
	handler lib.RequestHandler,
	authController controllers.AuthController,
	logger lib.Logger,
) AuthRoute {
	return AuthRoute{
		handler:        handler,
		logger:         logger,
		authController: authController,
	}
}
