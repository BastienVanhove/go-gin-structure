package contextStart

import (
	"context"
	blogContext "root/Context/Blog"
	global "root/Core/Global"
	model "root/Core/Model"
	"root/Core/envRead"
	"time"

	"github.com/gin-gonic/gin"
)

func Start() {

	//READ .env file
	fileEnvToRead := "context.env"
	envContext := envRead.Read(fileEnvToRead)

	env := ".env"
	envVariables := envRead.Read(env)

	//Context Initialisation
	ctx := context.TODO()

	//Context connexion: for timeout reponse
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	//GET connected to DB client
	DataBase := model.Connexion(ctx, envVariables["DB_NAME"], envVariables["DB_URI"])

	engine := gin.Default()

	//DEFINE Global struct
	Global := &global.Global{
		Name:       "AllContext",
		Engine:     engine,
		DataBase:   DataBase,
		AppContext: ctx,
	}

	Global.AddContext(blogContext.Init())
	Global.InitContexts(envContext)

	//engine.Run()
}
