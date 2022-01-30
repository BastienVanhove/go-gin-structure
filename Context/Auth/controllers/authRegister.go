package authController

import (
	authModel "root/Context/Auth/Models"
	auth "root/Core/Auth"
	global "root/Core/Global"
	utility "root/Core/Utility"

	"github.com/gin-gonic/gin"
)

func AuthRegister(global *global.Global, authGroup *gin.RouterGroup) {
	authGroup.GET("/register", func(c *gin.Context) {

		var login auth.Login
		if err := c.ShouldBind(&login); err != nil {
			panic(err)
		}
		if (login == auth.Login{}) {
			panic("Information manquant")
		}

		registerEntity := authModel.AuthRegisterEntity{
			DataBase:   global.DataBase,
			AppContext: global.AppContext,
		}

		password, _ := utility.HashPassword(login.Password)

		user := auth.User{
			Name:     login.UserName,
			Password: password,
		}

		registerEntity.Register(user)

		c.JSON(200, gin.H{
			"route": "/register",
		})
	})
}
