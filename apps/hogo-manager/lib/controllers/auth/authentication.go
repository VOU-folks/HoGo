package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateAuthenticationCode(ctx *gin.Context) {
	// ToDo: Create AuthenticationCode which can be exchanged with AccessToken
	ctx.JSON(http.StatusCreated, gin.H{
		"code": "HERE-GOES-AUTHENTICATION-CODE",
	})
}
