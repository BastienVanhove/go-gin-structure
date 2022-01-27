package contextManager

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Context struct {
	Name  string
	Start func(global *Global)
}

type Global struct {
	Name       string
	Engine     *gin.Engine
	DataBase   *mongo.Database
	AppContext context.Context
	Contexts   []Context
}

func (global *Global) AddContext(context *Context) {
	global.Contexts = append(global.Contexts, *context)
}

func (global *Global) InitContexts() {
	for _, context := range global.Contexts {
		//Verification ENV ici => if
		context.Start(global)
		//endif
	}
}
