package contextManage

import (
    "fmt"
	"root/Context/blog/controllers"
	"root/Context/home/controllers"

	contextManagerStruct "root/Core/ContextManager"

	"github.com/gin-gonic/gin"
)

func Start(mapEnv map[string]string) {

    fmt.Println(mapEnv)
    
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

	engine.Run()
}
