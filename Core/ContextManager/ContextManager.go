package contextManager

import (
	controler "root/Core/Controler"
	model "root/Core/Model"

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
