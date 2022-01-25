package contextManager

import (
	controler "testObjectRoot/Core/Controler"
	model "testObjectRoot/Core/Model"

	"github.com/gin-gonic/gin"
)

type ContextManager struct {
	contexts []Context
}

type Context struct {
	Name      string
	Engine    *gin.Engine
	Controler *controler.Controler
	Model     *model.Model
}
