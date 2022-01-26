package contextManager

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Context struct {
	Name  string
	Start func(global *Global)
}

//peut etre qu'il faudrait mettre l'engine dans context manager
type Global struct {
	Name     string
	Engine   *gin.Engine
	Contexts []Context
}

func (c *Global) AddContext(context *Context) {
	c.Contexts = append(c.Contexts, *context)
}
func (c *Global) StartAllContext(context *Context) {
	fmt.Println("start all context...")
}
