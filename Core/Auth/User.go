package auth

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Name       string             `bson:"name"`
	Password   string             `bson:"password"`
	Email      string             `bson:"email"`
	Provider   string             `json:"provider" bson:"provider"`
	IDProvider string             `json:"UserID" bson:"idprovider"`
}
