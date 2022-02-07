package auth

import (
	"context"
	authController "root/Core/Auth/Controller"
	authMiddleware "root/Core/Auth/Middleware"
	authModel "root/Core/Auth/Model"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Init(dataBase *mongo.Database, context context.Context, engine *gin.Engine) *jwt.GinJWTMiddleware {

	AuthEntity := &authModel.AuthEntity{
		DataBase:   dataBase,
		AppContext: context,
	}

	middleware := authMiddleware.Middleware(AuthEntity)

	authRoute := engine.Group("/auth")
	{
		authController.AuthRegister(AuthEntity, authRoute)

		authController.AuthLogin(middleware, authRoute)
		authController.AuthRefreshToken(middleware, authRoute)
		authController.AuthProvider(middleware, authRoute)
	}

	return middleware
}
