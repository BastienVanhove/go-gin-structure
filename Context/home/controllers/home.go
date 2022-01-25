package home

import (
	contextManager "root/Core/ContextManager"
    "github.com/gin-gonic/gin"
)

func AddTo(context *contextManager.Context) {
	Route(context)
}

func Route(context *contextManager.Context) {
	context.Engine.GET("/home", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"page": "home",
		})
	})
}
