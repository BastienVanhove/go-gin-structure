package contextManager

import (
	"fmt"
	controler "root/Core/Controler"
	model "root/Core/Model"

	"github.com/gin-gonic/gin"
)

type Context struct {
	Name      string
	Engine    *gin.Engine
	Controler *controler.Controler
	Model     *model.Model
}
//peut etre qu'il faudrait mettre l'engine dans context manager
type ContextManager struct {
	Name string
	Contexts map[string]Context
}

func (c ContextManager) AddContext(context func(cont *Context)){
	fmt.Println("new Context is : ", context)
}
