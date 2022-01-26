package contextManage

import (
    "fmt"
	"root/Context/blog/controllers"
	"root/Context/home/controllers"
	contextManagerStruct "root/Core/ContextManager"
	"github.com/gin-gonic/gin"
    "root/Core/envRead"
)

func Start() {

    //unactive section
    fileEnvToRead := "context.env"
    mapEnv := envRead.Read(fileEnvToRead)

    //read .env
    env := ".env"
    envVariable := envRead.Read(env)
    
    //ici faut faire un syst de 8080 de base que le port n'est pas def
    PORT := envVariable["PORT"]

    //im try here to put all context in struct contextManager

    allContextObj := &contextManagerStruct.ContextManager{
        Name: "AllContext",
        Contexts: make(map[string]contextManagerStruct.Context),
    }
    allContextObj.AddContext(home.AddTo)
    fmt.Println(allContextObj)

    //pour interagir avec les variables d'environnement
    allContext := make(map[string]func(context *contextManagerStruct.Context))
    
    //trouver un moyen d'automatiser avec les packages surrement
    allContext["home"] = home.AddTo
    allContext["blog"] = blog.AddTo

    //la logique d'env
    engine := gin.Default()

	Context := &contextManagerStruct.Context{Name: "Context", Engine: engine}

    for keyC, valueC := range allContext { 
        for keyEnv, valueEnv := range mapEnv{
            if(keyC == keyEnv && valueEnv == "false"){
                fmt.Println(keyEnv)
            }else{
                valueC(Context) 
            }
        }
    }

	engine.Run(":"+PORT)
}
