package contextManager

import (
	"fmt"

	controler "root/Core/Controler"

	"github.com/gin-gonic/gin"
)

type Context struct {
	Name      string
	Engine    *gin.Engine
	Controlers []controler.Controler
}

func (c *Context) AddController(controller controler.Controler){
	//il faut verifier si il n'existe pas deja
	c.Controlers = append(c.Controlers, controller)
}

func (c *Context) StartAllController() {
	fmt.Println("Starting all controllers")
}

//peut etre qu'il faudrait mettre l'engine dans context manager
type ContextManager struct {
	Name string
	Contexts map[string]Context
}

func (c ContextManager) AddContext(context *Context){
	fmt.Println("new Context is : ", context)
}
