package authModel

import (
	"errors"
	"fmt"
	authEntity "root/Core/Auth/Entity"
	model "root/Core/Model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthEntity model.Entity

func (e *AuthEntity) AuthLogin(email string) authEntity.User {
	var user authEntity.User
	collectionUser := e.DataBase.Collection("user")
	err := collectionUser.FindOne(e.AppContext, bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user
		}
		panic(err)
	}
	return user
}

func (e *AuthEntity) Register(user authEntity.User) (*mongo.InsertOneResult, error) {
	collectionUser := e.DataBase.Collection("user")

	var existingUser authEntity.User
	collectionUser.FindOne(e.AppContext, bson.D{{"email", user.Email}}).Decode(&existingUser)

	if (existingUser != authEntity.User{}) {
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
