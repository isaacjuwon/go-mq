package controller

import (
	"fusossafuoye.ng/app/response"
	user "fusossafuoye.ng/app/service"

	"github.com/gofiber/fiber/v2"
)

type DashboardController interface {
	Index(c *fiber.Ctx) error
}

type dashboardController struct {
	userService user.UserService
}

func NewDashboardController(userService user.UserService) DashboardController {
	return &dashboardController{
		userService: userService,
	}
}

func (ctrl *indexController) Login(c *fiber.Ctx) error {

	return response.SuccessDataResponse(c, fiber.StatusOK,
		token, "Login successful")
}
