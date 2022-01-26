package blogController

import (
	"fmt"
	contextManager "root/Core/ContextManager"

	"github.com/gin-gonic/gin"
)

func BlogComment(global *contextManager.Global) {
	global.Engine.GET("/blog/comment", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"route": "/blog/comment",
		})
	})
	fmt.Println("Route load")
}
