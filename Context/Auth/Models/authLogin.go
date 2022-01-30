package authModel

import (
	model "root/Core/Model"
	user "root/Core/User"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//IMPORTANT to create a struct on Entity => to add custom method for this model
type AuthLoginEntity model.Entity

func (e *AuthLoginEntity) Login(name string) user.User {
	var user user.User
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
