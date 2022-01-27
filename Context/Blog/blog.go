package blogContext

import (
	blogController "root/Context/blog/controllers"
	contextManager "root/Core/ContextManager"
)

func Init() *contextManager.Context {

	context := &contextManager.Context{
		Name: "blog",
		Start: func(global *contextManager.Global) {
			blogController.BlogComment(global)
			blogController.BlogUser(global)
		},
	}

	return context
}
