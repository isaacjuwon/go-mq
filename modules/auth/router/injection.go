package router

import (
	"fusossafuoye.ng/modules/auth/controller"
	"fusossafuoye.ng/app/repository"
	auth "fusossafuoye.ng/app/service"
	"fusossafuoye.ng/modules/auth/service"

)

type Injection struct {
	AuthController controller.AuthController
}

// Define Dependency Injection
func CallDependenciesInjection() Injection {


	userRepo := repository.NewUserRepository()
	userService := auth.NewUserService(userRepo)
	authService := service.NewAuthService(userRepo)
	authController := controller.NewAuthController(userService, authService)


	return Injection{
		AuthController: authController,
	}
}