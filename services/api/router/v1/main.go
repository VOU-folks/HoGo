package v1

import (
	"hogo/core/components"
)

func AttachV1Routes(app *components.Engine, router *components.RouterGroup) {
	usersRouter := router.Group("users")

	AttachUsersRoutes(app, usersRouter)
}
