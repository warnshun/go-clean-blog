package services

import "go.uber.org/fx"

// Module exports services present
var Module = fx.Options(
	fx.Provide(NewAuth),
	fx.Provide(NewUser),
	fx.Provide(NewPassword),
)
