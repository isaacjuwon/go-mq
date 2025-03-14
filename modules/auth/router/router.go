package router

import (
	"fusossafuoye.ng/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(a *fiber.App, c Injection) {
	auth := a.Group("/auth")

	authController := c.AuthController
	auth.Post("/register", middleware.DBTransactionHandler(), authController.RegisterUser)
	auth.Post("/login", authController.Login)

}