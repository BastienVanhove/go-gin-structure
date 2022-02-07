package blogController

import (
	authEnity "root/Core/Auth/Entity"
	authMiddleware "root/Core/Auth/Middleware"
	global "root/Core/Global"

	"github.com/gin-gonic/gin"
)

func BlogComment(global *global.Global) {

	r := global.Engine.Group("/blog/comment")

	r.Use(global.Auth.MiddlewareFunc())

	r.GET("/", func(c *gin.Context) {
		user, _ := c.Get(authMiddleware.IdentityKey)

		c.JSON(200, gin.H{
			"userID": user.(*authEnity.User).Email,
		})
	})
}
