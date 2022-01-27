package context

import (
	"context"
	blogContext "root/Context/Blog"
	contextManager "root/Core/ContextManager"
	model "root/Core/Model"
	"root/Core/envRead"
	"time"

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

	//Context Initialisation
	ctx := context.TODO()

	//Context connexion: for timeout reponse
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	//GET connected to DB client
	DataBase := model.Connexion(ctx, envVariable["DB_NAME"], envVariable["DB_URI"])

	engine := gin.Default()

	//DEFINE Global struct
	Global := &contextManager.Global{
		Name:       "AllContext",
		Engine:     engine,
		DataBase:   DataBase,
		AppContext: ctx,
	}

	Global.AddContext(blogContext.Init())
	Global.InitContexts()

	//engine.Run()
}
