package ContextManager

import "github.com/gin-gonic/gin"

type Context struct {
	Name   string
	Engine *gin.Engine
}
