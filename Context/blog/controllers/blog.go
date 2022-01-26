package blog

import (
	contextManager "root/Core/ContextManager"
)

func AddTo(context *contextManager.Context) *contextManager.Context{

	context.AddController( monControllerCommentaire(context) )
	context.AddController( blog1(context) )
	context.AddController( blog2(context) )

	return context
}
