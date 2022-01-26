package contextManage

import (
    "fmt"
	"root/Context/blog/controllers"
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

	Context := &contextManagerStruct.Context{
        Name: "Context", 
        Engine: engine,
    }
    allContextObj := &contextManagerStruct.ContextManager{
        Name: "AllContext",
        Contexts: make(map[string]contextManagerStruct.Context),
    }
    allContextObj.AddContext(blog.AddTo(Context))

    fmt.Println(allContextObj)

    

	//engine.Run(":"+PORT)
    
}
