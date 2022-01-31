package authModel

import (
	"errors"
	"fmt"
	auth "root/Core/Auth"
	model "root/Core/Model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//IMPORTANT to create a struct on Entity => to add custom method for this model
type AuthRegisterEntity model.Entity

func (e *AuthRegisterEntity) Register(user auth.User) (*mongo.InsertOneResult, error) {
	collectionUser := e.DataBase.Collection("user")

	var existingUser auth.User
	collectionUser.FindOne(e.AppContext, bson.D{{"email", user.Email}}).Decode(&existingUser)

	if (existingUser != auth.User{}) {
		fmt.Println("[Auth Register] Email deja enregistré")
		return &mongo.InsertOneResult{}, errors.New("email deja enregistré")
	}

	fmt.Println("[Auth Register] Création du nouvel utilisateur")
	result, err := collectionUser.InsertOne(e.AppContext, user)
	if err != nil {
		panic(err)
	}
	return result, nil

}
