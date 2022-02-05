package auth

import (
	utility "root/Core/Utility"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/gin-gonic/gin"
)

var IdentityKey = "id"

//TODO : create a package Auth
//TODO : create cookie with (Secure, HttpOnly, SameSite) attributs

func Middleware(dataBase *mongo.Database) *jwt.GinJWTMiddleware {

	AuthEntity := AuthEntity{
		DataBase: dataBase,
	}

	authMiddleware, _ := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if user, ok := data.(*User); ok {
				return jwt.MapClaims{
					IdentityKey: user.ID,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			//TODO: 
			claims := jwt.ExtractClaims(c)
			objectID, _ := primitive.ObjectIDFromHex(claims[IdentityKey].(string))
			return &User{
				ID: objectID,
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
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*User); ok {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",

		TokenHeadName: "Bearer",

		TimeFunc: time.Now,
	})

	return authMiddleware
}
