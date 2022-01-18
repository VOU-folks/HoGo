package v1

import (
	"hogo/core/components"
	"hogo/services/api/handlers/v1/users"
)

func AttachUsersRoutes(app *components.Engine, router *components.RouterGroup) {
	router.GET("", users.ListUsers)
	router.POST("", users.CreateUser)
	router.GET(":id", users.GetUser)
	router.PUT(":id", users.UpdateUser)
	router.PATCH(":id", users.UpdateUser)

	router.PATCH(":id/active", users.ActivateUser)
	router.PATCH(":id/deactivate", users.DeactivateUser)

	router.DELETE(":id", users.DeleteUser)
	router.PATCH(":id/undelete", users.UndeleteUser)
}
