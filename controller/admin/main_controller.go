package admin

import "kang-sayur-backend/infrastructure/injector"

type mainController struct {
	injector injector.InjectorProvider
}

func MainController(injector injector.InjectorProvider) mainController {
	return mainController{injector: injector}
}
