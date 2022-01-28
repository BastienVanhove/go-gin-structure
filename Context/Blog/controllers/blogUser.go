package blogController

import (
	global "root/Core/Global"

	"github.com/gin-gonic/gin"
)

func BlogUser(global *global.Global) {
	global.Engine.GET("/blog/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"route": "/blog/user",
		})
	})
}
