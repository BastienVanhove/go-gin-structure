package blogController

import (
	global "root/Core/Global"

	"github.com/gin-gonic/gin"
)

func BlogComment(global *global.Global) {
	global.Engine.GET("/blog/comment", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"route": "/blog/comment",
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
