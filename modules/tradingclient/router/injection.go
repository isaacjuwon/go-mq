package router

import (
	"fusossafuoye.ng/modules/tradingclient/controller"
	"fusossafuoye.ng/app/repository"
	user "fusossafuoye.ng/app/service"

)

type Injection struct {
	DashboardController controller.DashboardController
}

// Define Dependency Injection
func CallDependenciesInjection() Injection {


	userRepo := repository.NewUserRepository()
	userService :=dashboard.NewUserService(userRepo)
	dashboardController := controller.NewDashboardController(userService)


	return Injection{
		DashboardController: dashboardController,
	}
}