package context

import (
	"context"
	blogContext "root/Context/Blog"
	contextManager "root/Core/ContextManager"
	model "root/Core/Model"
	"root/Core/envRead"

	"github.com/gin-gonic/gin"
)

func Start() {

	/*
	   //unactive section
	   fileEnvToRead := "context.env"
	   mapEnv := envRead.Read(fileEnvToRead)
	*/

	//READ .env file
	env := ".env"
	envVariable := envRead.Read(env)

	//Context connexion: for timeout reponse
	ctx := context.TODO()

	//GET connected to DB client
	DataBase := model.Connexion(envVariable["DB_NAME"], envVariable["DB_URI"], ctx)

	engine := gin.Default()

	//DEFINE Global struct
	allContextObj := &contextManager.Global{
		Name:       "AllContext",
		Engine:     engine,
		DataBase:   DataBase,
		AppContext: ctx,
	}

	allContextObj.AddContext(blogContext.Init())

	InitContexts(allContextObj)

	//engine.Run()
}

func InitContexts(global *contextManager.Global) {
	for _, context := range global.Contexts {
		//Verification ENV ici => if
		context.Start(global)
		//endif
	}
}
