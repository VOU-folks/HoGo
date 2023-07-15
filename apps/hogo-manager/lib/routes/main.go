package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Attach(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, &gin.H{
			"type":    "service",
			"name":    "hogo-manager",
			"version": "1.0.0",
		})
	})
	AttachAuth(r.Group("/auth"))
}
