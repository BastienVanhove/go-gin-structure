package blog

import (
	"fmt"
	contextManager "root/Core/ContextManager"

)

func AddTo(context *contextManager.Context) *contextManager.Context{

	context.AddController( monControllerCommentaire(context) )
	context.AddController( blog1(context) )

	fmt.Println( context )

	return context
}
