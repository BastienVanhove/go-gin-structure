package auth

import (
	model "root/Core/Model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthEntity model.Entity

func (e *AuthEntity) AuthLogin(name string) User {
	var user User
	collectionUser := e.DataBase.Collection("user")
	err := collectionUser.FindOne(e.AppContext, bson.D{{"name", name}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user
		}
		panic(err)
	}
	return user
}
