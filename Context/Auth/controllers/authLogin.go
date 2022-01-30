package authController

import (
	global "root/Core/Global"

	"github.com/gin-gonic/gin"
)

func AuthLogin(global *global.Global, authGroup *gin.RouterGroup) {
	authGroup.GET("/login", global.Auth.LoginHandler)
}
