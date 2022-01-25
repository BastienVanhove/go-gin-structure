package home

import (
	contextManager "testObjectRoot/Core/ContextManager"

	"github.com/gin-gonic/gin"
)

func New(context *contextManager.Context) {
	Route(context)
}

func Route(context *contextManager.Context) {
	context.Engine.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"page": "home",
		})
	})
}
