package contextManager

import (
	controler "testObjectRoot/Core/Controler"

	"github.com/gin-gonic/gin"
)

type ContextManager struct {
	contexts []Context
}

type Context struct {
	Name      string
	Engine    *gin.Engine
	Controler *controler.Controler
}
