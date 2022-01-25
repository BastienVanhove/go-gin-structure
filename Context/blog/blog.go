package blog

import (
	"testObjectRoot/Core/ContextManager"

	"github.com/gin-gonic/gin"
)

func New(context *ContextManager.Context) {
	Route(context)
}

func Route(context *ContextManager.Context) {
	context.Engine.GET("/blog", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"page": "blog",
		})
	})
}
