package bootstrap

import (
	"go-clean-blog/api/controllers"
	"go-clean-blog/api/middlewares"
	"go-clean-blog/api/routes"
	"go-clean-blog/lib"
	"go-clean-blog/repository"
	"go-clean-blog/services"

	"go.uber.org/fx"
)

var CommonModules = fx.Options(
	controllers.Module,
	routes.Module,
	lib.Module,
	services.Module,
	middlewares.Module,
	repository.Module,
)
