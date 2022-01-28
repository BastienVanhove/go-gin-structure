package global

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type Global struct {
	Name       string
	Engine     *gin.Engine
	DataBase   *mongo.Database
	AppContext context.Context
	Contexts   []ContextController
}

func (global *Global) AddContext(context *ContextController) {
	global.Contexts = append(global.Contexts, *context)
}

func (global *Global) InitContexts(envContext map[string]string) {
	for _, context := range global.Contexts {
		if envContext[context.Name] != "false" {
			context.Start(global)
		}
	}
}

/*
	Problem of "cycle import"
	=> Move into Core/ContextController needs some structure edits
*/
type ContextController struct {
	Name  string
	Start func(global *Global)
}
