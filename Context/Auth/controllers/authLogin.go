package authController

import (
	"fmt"
	authModel "root/Context/Auth/Models"
	global "root/Core/Global"
	user "root/Core/User"
	utility "root/Core/Utility"

	"github.com/gin-gonic/gin"
)

func AuthLogin(global *global.Global, auth *gin.RouterGroup) {
	auth.GET("/login", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"route": "/blog/comment",
		})
	})

	loginEntity := authModel.AuthLoginEntity{
		DataBase:   global.DataBase,
		AppContext: global.AppContext,
	}

	name := "Thomas"
	password := "passwords"

	u := loginEntity.Login(name)

	if (user.User{} == u) {
		fmt.Println("[Login] Utilisateur introuvable")
	} else {
		fmt.Println("[Login] Utilisateur trouvé")
		if utility.CheckPasswordHash(password, u.Password) {
			fmt.Println("[Login] Connexion accepté")
			fmt.Print(u)
		} else {
			fmt.Println("[Login] Mot de passe incorrect")
		}
	}
}
