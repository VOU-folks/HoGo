package routes

import (
	"github.com/gin-gonic/gin"
	"hogo/apps/hogo-manager/lib/controllers/auth"
)

func AuthRouting(r *gin.RouterGroup) {
	r.POST("/code", auth.CreateAuthenticationCode) // username/password -> auth code
	r.POST("/session", auth.CreateSession)         // auth code -> access token
	r.DELETE("/session", auth.DestroySession)
}
