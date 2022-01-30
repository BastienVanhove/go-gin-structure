package authModel

import (
	model "root/Core/Model"
	user "root/Core/User"

	"go.mongodb.org/mongo-driver/mongo"
)

//IMPORTANT to create a struct on Entity => to add custom method for this model
type AuthRegisterEntity model.Entity

func (e *AuthRegisterEntity) Register(user user.User) *mongo.InsertOneResult {
	collectionUser := e.DataBase.Collection("user")
	result, err := collectionUser.InsertOne(e.AppContext, user)
	if err != nil {
		panic(err)
	}
	return result
}
