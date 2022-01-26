package main

import (
    "fmt"
    contextManage "root/Context/contextManage"
   "root/Core/envRead"
)

func main() {
    //unactive section
    fileEnvToRead := "context.env"
    contextEnabled := envRead.Read(fileEnvToRead)
    contextManage.Start(contextEnabled)

    //read .env
    env := ".env"
    envVariable := envRead.Read(env)
    fmt.Println(envVariable)
}
