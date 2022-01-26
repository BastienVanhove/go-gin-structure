package blogController

import (
	"fmt"
	blogModel "root/Context/Blog/Models"
	contextManager "root/Core/ContextManager"

	"github.com/gin-gonic/gin"
)

func BlogComment(global *contextManager.Global) {
	global.Engine.GET("/blog/comment", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"route": "/blog/comment",
		})
	})

	commentEntity := blogModel.BlogCommentEntity{
		DataBase: global.DataBase,
	}

	commentEntity.GetComment(1)
	fmt.Println(commentEntity.Message)
}
