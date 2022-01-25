package main

import (
	"testObjectRoot/Context/blog"
	"testObjectRoot/Context/home"
	"testObjectRoot/Manager/ContextManager"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	context := &ContextManager.Context{Name: "Context", Engine: engine}
	blog.New(context)
	home.New(context)

	engine.Run()
}
