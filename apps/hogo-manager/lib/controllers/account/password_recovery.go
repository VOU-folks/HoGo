package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePasswordRecoveryRequest(ctx *gin.Context) {
	// ToDo: Handle user credential to send PasswordRecoveryRequest to email
	ctx.JSON(http.StatusOK, gin.H{})
}

func ConfirmPasswordRecoveryRequest(ctx *gin.Context) {
	// ToDo: Handle PasswordRecoveryRequest data and new password to recover access
	// ToDo: Save history of PasswordRecoverRequest and passwords to track user behaviour for security hardening purposes
	ctx.JSON(http.StatusOK, gin.H{})
}
