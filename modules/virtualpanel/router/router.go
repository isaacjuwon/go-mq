package router

import (
	"fusossafuoye.ng/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(a *fiber.App, c Injection) {
	route := a.Group("/virtualpanel")
	users := route.Group("/users")

	userController := c.UserController

	users.Get("", middleware.Protected(), userController.GetUsers)
	users.Get("/:id", middleware.Protected(), userController.GetUserByID)
	users.Post("", middleware.DBTransactionHandler(), userController.CreateUser)
	users.Put("/:id", middleware.DBTransactionHandler(), userController.UpdateUser)
	users.Delete("/:id", middleware.DBTransactionHandler(), userController.DeleteUser)

}
