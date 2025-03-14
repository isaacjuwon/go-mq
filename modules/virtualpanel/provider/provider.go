package provider

import (
	"fusossafuoye.ng/modules/virtualpanel/router"
	"github.com/gofiber/fiber/v2"
)

func SetupProvider(a *fiber.App) {
	// Dependencies Injection
	injection := router.CallDependenciesInjection()

	// Routes.
	router.SetupRoutes(a, injection)
}
 