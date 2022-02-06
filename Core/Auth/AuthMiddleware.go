package auth

import (
	"fmt"
	utility "root/Core/Utility"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

var (
	key            = []byte("secret key")
	IdentityKey    = "id"
	timeout        = time.Hour
	maxRefresh     = time.Hour
	sendCookie     = true
	secureCookie   = false
	cookieHTTPOnly = true
)

//TODO : create a package Auth
//TODO : create cookie with (Secure, HttpOnly, SameSite) attributs

type AuthManager struct {
	Classic *jwt.GinJWTMiddleware
	OAuth   *jwt.GinJWTMiddleware
}

func Middleware(dataBase *mongo.Database) *AuthManager {
	return &AuthManager{Classic: MiddlewareClassicAuth(dataBase), OAuth: MiddlewareClassicAuth(dataBase)}
}

func MiddlewareClassicAuth(dataBase *mongo.Database) *jwt.GinJWTMiddleware {

	AuthEntity := AuthEntity{
		DataBase: dataBase,
	}

	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         key,
		Timeout:     timeout,
		MaxRefresh:  maxRefresh,
		IdentityKey: IdentityKey,

		SendCookie:     sendCookie,
		SecureCookie:   secureCookie, // true when in prod with https
		CookieHTTPOnly: cookieHTTPOnly,

		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*User); ok {
				fmt.Println(user)
				return jwt.MapClaims{
					IdentityKey: user.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Email: claims[IdentityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals Login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user := AuthEntity.AuthLogin(loginVals.UserName)
			if (User{} != user && utility.CheckPasswordHash(loginVals.Password, user.Password)) {
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: authorizator,
		Unauthorized: unauthorized,
		TokenLookup:  "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})

	return authMiddleware
}

func authorizator(data interface{}, c *gin.Context) bool {
	if _, ok := data.(*User); ok {
		return true
	}
	return false
}

func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
