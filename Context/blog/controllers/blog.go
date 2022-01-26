package blog

import (
	"fmt"
	contextManager "root/Core/ContextManager"
	"github.com/gin-gonic/gin"
)

func AddTo(context *contextManager.Context) {
	Route(context)
}

func Route(context *contextManager.Context) {
	fmt.Println(context)
	context.Engine.GET("/blog", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"page": "blog",
		})
	})
	context.Engine.GET("/blog1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"page": "blogNumero1",
		})
	})
}
