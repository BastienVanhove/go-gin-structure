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

func (c *Global) AddContext(context *Context) {
	c.Contexts = append(c.Contexts, *context)
}
