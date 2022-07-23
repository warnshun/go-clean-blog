package routes

import (
	"go-clean-blog/api/controllers"
	"go-clean-blog/api/middlewares"
	"go-clean-blog/lib"
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
		api.GET("/post/:id", r.postController.GetPost)
		api.GET("/post", r.postController.GetAllPosts)
		api.POST("/post", r.postController.AddPost)

		api.POST("/post/:id/like", r.postController.SwitchLikePost)

		api.GET("/liked_post", r.postController.GetAllLikedPosts)
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
