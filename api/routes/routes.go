package routes

import "go.uber.org/fx"

// Module exports dependency to container
var Module = fx.Options(
	fx.Provide(NewAuthRoute),
	fx.Provide(NewUserRoute),
	fx.Provide(NewPostRoute),
	fx.Provide(NewRoutes),
)

// Routes contains multiple routes
type Routes []Route

// Route interface
type Route interface {
	Setup()
}

// NewRoutes sets up routes
func NewRoutes(
	authRoutes AuthRoute,
	userRoutes UserRoute,
	postRoutes PostRoute,
) Routes {
	return Routes{
		authRoutes,
		userRoutes,
		postRoutes,
	}
}

// Setup all the route
func (r Routes) Setup() {
	for _, route := range r {
		route.Setup()
	}
}
