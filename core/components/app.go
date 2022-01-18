package components

import (
	"github.com/gin-gonic/gin"
)

type (
	Context     = gin.Context
	Engine      = gin.Engine
	Router      = gin.IRouter
	RouterGroup = gin.RouterGroup
)

func NewHttpAppInstance() *Engine {
	return gin.New()
}
