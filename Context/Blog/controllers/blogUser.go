package blogController

import (
	"fmt"
	contextManager "root/Core/ContextManager"

	"github.com/gin-gonic/gin"
)

func BlogUser(global *contextManager.Global) {
	global.Engine.GET("/blog/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"route": "/blog/user",
		})
	})
	fmt.Println("Route load")
}
