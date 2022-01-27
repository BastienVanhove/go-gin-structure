package blogController

import (
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
		DataBase:   global.DataBase,
		AppContext: global.AppContext,
	}

	myComment := blogModel.Comment{
		User:    "T",
		Message: "Heyy",
	}

	yourComment := blogModel.Comment{
		User:    "B",
		Message: "Awesome !!",
	}

	commentEntity.CreateComment(myComment)
	commentEntity.CreateComment(yourComment)
}
