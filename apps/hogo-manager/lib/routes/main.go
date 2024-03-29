package routes

import (
	"github.com/gin-gonic/gin"
	"hogo/lib/core/globals"
	"net/http"
)

func Attach(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, globals.AppInfo)
	})

	AuthRouting(r.Group("/auth"))
	AccountRouting(r.Group("/account"))
}
