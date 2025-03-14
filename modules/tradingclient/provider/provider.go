package provider

import (
	"fusossafuoye.ng/modules/auth/router"
	"github.com/gofiber/fiber/v2"
)

func SetupProvider(a *fiber.App) {
	// Dependencies tradingclient Injection
	injection := router.CallDependenciesInjection()

	// Routes.
	router.SetupRoutes(a, injection)
}
