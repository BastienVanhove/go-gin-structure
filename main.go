package main

import (
	"testObjectRoot/Context/blog"
	"testObjectRoot/Context/home"
	contextManager "testObjectRoot/Core/ContextManager"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	context := &contextManager.Context{Name: "Context", Engine: engine}
	blog.New(context)
	home.New(context)

	engine.Run()
}
