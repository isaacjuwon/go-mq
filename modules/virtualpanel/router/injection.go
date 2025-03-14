package router

import (
	"fusossafuoye.ng/modules/virtualpanel/controller"
	"fusossafuoye.ng/modules/virtualpanel/repository"
	"fusossafuoye.ng/modules/virtualpanel/service"
)

type Injection struct {
	UserController controller.UserController
}

// Define Dependency Injection
func CallDependenciesInjection() Injection {
	
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)


	return Injection{
		UserController: userController,
	}
}