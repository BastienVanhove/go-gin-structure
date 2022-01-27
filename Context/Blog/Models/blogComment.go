package blogModel

import (
	"fmt"
	model "root/Core/Model"
)

type BlogCommentEntity struct {
	DataBase *model.DataBaseSession
	User     string
	Message  string
}

func (e *BlogCommentEntity) CreateComment(comment string) {
	fmt.Println("Connexion : ", e.DataBase.Host, e.DataBase.Password)
	fmt.Println("[Model blogComment] Success created comment => ", comment)
}

func (e *BlogCommentEntity) GetComment(id int) string {
	message := "Un message recuperé dans la DB"
	e.Message = message
	return message
}
