package context

import (
	"fmt"
	blogContext "root/Context/Blog"
	contextManager "root/Core/ContextManager"
	model "root/Core/Model"

	"github.com/gin-gonic/gin"
	//"root/Core/envRead"
)

func Start() {

	/*
	   //unactive section
	   fileEnvToRead := "context.env"
	   mapEnv := envRead.Read(fileEnvToRead)

	   //read .env
	   env := ".env"
	   envVariable := envRead.Read(env)

	   //ici faut faire un syst de 8080 de base que le port n'est pas def
	   PORT := envVariable["PORT"]
	*/

	engine := gin.Default()

	mongodb := &model.DataBaseSession{
		Host:     "localhost",
		Password: "***",
	}

	//entityManager := &model.EntityManager{}

	allContextObj := &contextManager.Global{
		Name:   "AllContext",
		Engine: engine,
		//EntityManager: entityManager,
		DataBase: mongodb,
	}

	allContextObj.AddContext(blogContext.Init())

	Route(allContextObj)

	fmt.Println(allContextObj)

	//engine.Run()

}

func Route(global *contextManager.Global) {
	for _, context := range global.Contexts {
		//Verification ENV ici => if
		fmt.Println("ROUTE")
		context.Start(global)
		//endif
	}
}
