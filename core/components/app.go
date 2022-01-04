package components

import (
    "github.com/gin-gonic/gin"
)

type (
    Context = gin.Context
    Engine  = gin.Engine
)

func NewHttpAppInstance() *Engine {
    return gin.New()
}