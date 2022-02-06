package authController

import (
	"encoding/json"
	"net/http"
	authModel "root/Context/Auth/Models"
	auth "root/Core/Auth"
	global "root/Core/Global"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

func AuthLogin(global *global.Global, authGroup *gin.RouterGroup) {

	authGroup.GET("/login", global.Auth.LoginHandler)
	authGroup.GET("/refresh_token", global.Auth.RefreshHandler)

	//authGroup.GET("/login", func(c *gin.Context) {
	//	Auth := global.Auth
	//	data, err := Auth.Authenticator(c)
	//	if err != nil {
	//		panic("NOP")
	//	}
	//	fmt.Println("Data login : ", data)
	//})

	key := "Secret-session-key" // Replace with your SESSION_SECRET or similar
	maxAge := 86400 * 30        // 30 days
	isProd := false             // Set to true when serving over https

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true // HttpOnly should always be enabled
	store.Options.Secure = isProd
	gothic.Store = store

	goth.UseProviders(
		google.New("463008552495-v343bl24nekr11utnftg9u6neuuoge45.apps.googleusercontent.com", "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", "http://localhost:3000/auth/google/callback", "email", "profile"),
	)

	//TODO: Put in Auth
	authGroup.GET("/:provider", func(c *gin.Context) {
		provider := c.Param("provider")
		q := c.Request.URL.Query()
		q.Add("provider", provider)
		c.Request.URL.RawQuery = q.Encode()
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})
	//authGroup.GET("/test", func(c *gin.Context) {
	//	r, _ := c.Cookie("_gothic_session")
	//	fmt.Println(r)
	//})

	authGroup.GET("/:provider/callback", func(c *gin.Context) {
		provider := c.Param("provider")
		q := c.Request.URL.Query()
		q.Add("provider", provider)
		c.Request.URL.RawQuery = q.Encode()
		data, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		var user auth.User
		res, err := json.Marshal(data)
		err = json.Unmarshal([]byte(res), &user)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		user.UseProvider = true

		registerEntity := authModel.AuthRegisterEntity{
			DataBase:   global.DataBase,
			AppContext: global.AppContext,
		}

		/*TODO:
		If exist :
			Login
		else
			Register
		*/

		_, err = registerEntity.Register(user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Failed",
				"error":   err.Error(),
			})
			return
		}

		tokenString, _, err := global.Auth.TokenGenerator(&user)
		if err != nil {
			panic(err)
		}
		expireCookie := global.Auth.TimeFunc().Add(global.Auth.CookieMaxAge)
		maxage := int(expireCookie.Unix() - global.Auth.TimeFunc().Unix())

		if global.Auth.SendCookie {

			if global.Auth.CookieSameSite != 0 {
				c.SetSameSite(global.Auth.CookieSameSite)
			}

			c.SetCookie(
				global.Auth.CookieName,
				tokenString,
				maxage,
				"/",
				global.Auth.CookieDomain,
				global.Auth.SecureCookie,
				global.Auth.CookieHTTPOnly,
			)
		}

		c.JSON(200, gin.H{
			"message": "success",
			"user":    user,
		})

	})
}
