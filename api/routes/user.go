package routes

import (
	"github.com/dipeshdulal/clean-gin/api/controllers"
	"github.com/dipeshdulal/clean-gin/api/middlewares"
	"github.com/dipeshdulal/clean-gin/lib"
)

// UserRoute struct
type UserRoute struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	userController controllers.UserController
	authMiddleware middlewares.JWTAuthMiddleware
}

// Setup user routes
func (r UserRoute) Setup() {
	r.logger.Info("Setting up user routes")
	api := r.handler.Gin.Group("/api").Use(r.authMiddleware.Handler())
	{
		// user
		api.GET("/user", r.userController.GetAllUsers)
		api.GET("/user/:id", r.userController.GetOneUser)
	}
}

// NewUserRoute creates new user controller
func NewUserRoute(
	logger lib.Logger,
	handler lib.RequestHandler,
	userController controllers.UserController,
	authMiddleware middlewares.JWTAuthMiddleware,
) UserRoute {
	return UserRoute{
		handler:        handler,
		logger:         logger,
		userController: userController,
		authMiddleware: authMiddleware,
	}
}
