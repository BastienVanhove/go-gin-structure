package authContext

import (
	authController "root/Context/Auth/controllers"
	global "root/Core/Global"
)

func Init() *global.ContextController {

	context := &global.ContextController{
		Name: "auth",
		Start: func(global *global.Global) {
			auth := global.Engine.Group("/auth")
			{
				authController.AuthLogin(global, auth)
				authController.AuthRegister(global, auth)
			}
		},
	}

	return context
}
