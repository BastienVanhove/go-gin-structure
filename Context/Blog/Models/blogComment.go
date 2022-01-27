package blogModel

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type BlogCommentEntity struct {
	DataBase   *mongo.Database
	AppContext context.Context
}

type Comment struct {
	User    string
	Message string
}

var collectionName string = "blogComment"
var collectionComment *mongo.Collection

func (e *BlogCommentEntity) CreateComment(comment Comment) *mongo.InsertOneResult {
	collectionComment := e.DataBase.Collection(collectionName)
	result, err := collectionComment.InsertOne(e.AppContext, comment)
	if err != nil {
		panic(err)
	}
	return result
}

func (e *BlogCommentEntity) GetComment(user string) {
}
