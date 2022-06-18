package services

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewAuthService),
	fx.Provide(NewUserService),
	fx.Provide(NewPostService),
)
