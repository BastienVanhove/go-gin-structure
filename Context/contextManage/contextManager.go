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
    fmt.Println(PORT)
    
    //pour interagir avec les variables d'environnement
    allContext := make(map[string]func(context *contextManagerStruct.Context))
    
    //trouver un moyen d'automatiser avec les packages surrement
    allContext["home"] = home.AddTo
    allContext["blog"] = blog.AddTo

    engine := gin.Default()

	contextM := &contextManagerStruct.Context{Name: "ContextM", Engine: engine}
    for keyC, valueC := range allContext { 
        for keyEnv, valueEnv := range mapEnv{
            if(keyC == keyEnv && valueEnv == "false"){
                fmt.Println(keyEnv)
            }else{
                valueC(contextM) 
            }
        }
    }

	engine.Run(":"+PORT)
}
