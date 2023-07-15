package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateSession(ctx *gin.Context) {
	// ToDo: Handle AuthenticationCode and create AccessToken
	ctx.JSON(http.StatusCreated, gin.H{
		"accessToken": "HERE-GOES-ACCESS-TOKEN",
	})
}

func DestroySession(ctx *gin.Context) {
	// ToDo: Delete AccessToken
	ctx.JSON(http.StatusOK, gin.H{})
}
