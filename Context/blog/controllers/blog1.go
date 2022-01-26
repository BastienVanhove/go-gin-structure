package blog 

import (
	controler "root/Core/Controler"
	contextManager "root/Core/ContextManager"
	"github.com/gin-gonic/gin"
)

func blog1(context *contextManager.Context) controler.Controler{
	var data func(string)
	data = func(route string){
		context.Engine.GET(route, func(c *gin.Context) {
			c.JSON(200, gin.H{
				"data": "des datas",
			})
		})
	}
	controller := controler.Controler{
		Route: "/blog1",
		Start: data,
    }
	return controller	
}