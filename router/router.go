package router

import (
	"fusossafuoye.ng/router/api"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app          *fiber.App
	healthRouter *api.HealthRouter
	
}

func New(app *fiber.App) *Router {
	return &Router{
		app:          app,
		healthRouter: api.NewHealthRouter(app),
	}
}

func Setup(app *fiber.App) {
	router := New(app)
	app.Stack()

	// Setup API routes with rate limiter
	apiRoute := app.Group("/api")

	// Setup individual route groups
	router.healthRouter.Setup(apiRoute)
}
