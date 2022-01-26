package context

import (
	"fmt"
	blog "root/Context/blog/controllers"
	contextManagerStruct "root/Core/ContextManager"

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

	allContextObj := &contextManagerStruct.Global{
		Name: "AllContext",
	}

	Context := &contextManagerStruct.Context{
		Name:   "Blog",
		Engine: engine,
	}
	allContextObj.AddContext(blog.AddTo(Context))

	fmt.Println(allContextObj)

	//engine.Run(":"+PORT)

}
