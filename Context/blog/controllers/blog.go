package blog

import (
	"fmt"
	contextManager "root/Core/ContextManager"
	"github.com/gin-gonic/gin"
)

func AddTo(context *contextManager.Context) {
	Route(context)
	//rename cette function New
	//renvoyer un objet de type Context qui contient tous les objet de 
	//type controllers dans un arr 
	//l'objet context a aussi une methods pour loop start tous les controllers
}

func Route(context *contextManager.Context) {
	fmt.Println(context)
	context.Engine.GET("/blog", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"page": "blog",
		})
	})
	context.Engine.GET("/blog1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"page": "blogNumero1",
		})
	})
	//ici on ajoutera les controlleurs dans la map controllers du struct Context
	//chaque objet controller aura un nom et une func Start(route string)
	//context.controllers = append(context.controllers, leController)
	//donc Context aura une methods pour ajouter des objet controller dans ca liste
	//en plus de la methods qui permet de tous les start (apres en fonxtion des variable d'environnement)
	//donc on aura un truc du style
	//dans un fichier separer
		//func getPage1(route string){
			//gin logic
		//}
		//controller1{name: "controller1", getPage1} 
	//dans la func Route:
	//context.addRoute(controller1)

}
