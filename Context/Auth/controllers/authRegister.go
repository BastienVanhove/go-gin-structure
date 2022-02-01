package authController

import (
	"net/http"
	authModel "root/Context/Auth/Models"
	auth "root/Core/Auth"
	global "root/Core/Global"
	utility "root/Core/Utility"

	"github.com/gin-gonic/gin"
)

func AuthRegister(global *global.Global, authGroup *gin.RouterGroup) {
	authGroup.GET("/register", func(c *gin.Context) {

		var register auth.Register
		if err := c.ShouldBind(&register); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"errors": utility.BeautifulError(err)})
			return
		}

		registerEntity := authModel.AuthRegisterEntity{
			DataBase:   global.DataBase,
			AppContext: global.AppContext,
		}

		password, _ := utility.HashPassword(register.Password)

		user := auth.User{
			Name:     register.UserName,
			Email:    register.Email,
			Password: password,
		}

		id, err := registerEntity.Register(user)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"error":   err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"message":  "success",
			"username": user.Name,
			"email":    user.Email,
			"password": user.Password,
			"ID":       id,
		})

	})
}
