package routes

import (
	"github.com/gin-gonic/gin"
	"hogo/apps/hogo-manager/lib/controllers/account"
)

func AccountRouting(r *gin.RouterGroup) {
	r.POST("/password-recovery", account.CreatePasswordRecoveryRequest) // username or email -> recovery request
	r.PUT("/password-recovery", account.ConfirmPasswordRecoveryRequest) // recovery request -> new password
}
