package routes

import (
	"github.com/dipeshdulal/clean-gin/api/controllers"
	"github.com/dipeshdulal/clean-gin/api/middlewares"
	"github.com/dipeshdulal/clean-gin/lib"
)

type PostRoute struct {
	logger         lib.Logger
	handler        lib.RequestHandler
	authMiddleware middlewares.JWTAuthMiddleware
	postController controllers.PostController
}

func (r PostRoute) Setup() {
	r.logger.Info("Setting up post routes")
	api := r.handler.Gin.Group("/api").Use(r.authMiddleware.Handler())
	{
		// post
		api.POST("/post", r.postController.AddPost)
	}
}

func NewPostRoute(
	logger lib.Logger,
	handler lib.RequestHandler,
	authMiddleware middlewares.JWTAuthMiddleware,
	postController controllers.PostController,
) PostRoute {
	return PostRoute{
		handler:        handler,
		logger:         logger,
		authMiddleware: authMiddleware,
		postController: postController,
	}
}