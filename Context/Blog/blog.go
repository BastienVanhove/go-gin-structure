package blogContext

import (
	blogController "root/Context/blog/controllers"
	global "root/Core/Global"
)

func Init() *global.ContextController {

	context := &global.ContextController{
		Name: "blog",
		Start: func(global *global.Global) {
			blogController.BlogComment(global)
			blogController.BlogUser(global)
		},
	}

	return context
}
