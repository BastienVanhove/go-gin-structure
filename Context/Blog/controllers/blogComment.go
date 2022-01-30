package blogController

import (
	"fmt"
	auth "root/Core/Auth"
	global "root/Core/Global"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func BlogComment(global *global.Global) {

	r := global.Engine.Group("/blog/comment")

	r.Use(global.Auth.MiddlewareFunc())

	r.GET("/", func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)

		user, _ := c.Get("id")

		fmt.Println(claims)

		c.JSON(200, gin.H{
			"user": user.(*auth.User),
		})
	})

	//commentEntity := blogModel.BlogCommentEntity{
	//	DataBase:   global.DataBase,
	//	AppContext: global.AppContext,
	//}

	//myComment := blogModel.Comment{
	//	User:    "T",
	//	Message: "Heyy",
	//}

	//commentEntity.CreateComment(myComment)
}
