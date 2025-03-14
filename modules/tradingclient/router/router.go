package router

import (
	"fusossafuoye.ng/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(a *fiber.App, c Injection) {
	client := a.Group("/trading-client")

	dashboardController := c.DashboardController
	
	dashboard.Post("", dashboardController.Index)

}