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

	//context.AddController( blog1(context) )
	//context.AddController( blog2(context) )

	return context
}
