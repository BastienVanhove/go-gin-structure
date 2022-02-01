package contextStart

import (
	"context"
	"reflect"
	authContext "root/Context/Auth"
	blogContext "root/Context/Blog"
	auth "root/Core/Auth"
	global "root/Core/Global"
	model "root/Core/Model"
	"root/Core/envRead"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
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
	//ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	//defer cancel()

	//GET connected to DB client
	DataBase := model.Connexion(ctx, envVariables["DB_NAME"], envVariables["DB_URI"])

	engine := gin.Default()

	//change f.Field() validator.ValidationErrors => return the `json tag` instead of `name` in struct
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	//DEFINE Global struct
	Global := &global.Global{
		Engine:     engine,
		DataBase:   DataBase,
		Auth:       auth.Middleware(DataBase),
		AppContext: ctx,
	}

	Global.AddContext(blogContext.Init())
	Global.AddContext(authContext.Init())
	Global.InitContexts(envContext)

	engine.Run()
}
