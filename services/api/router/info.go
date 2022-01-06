package router

import (
	"hogo/core/components"
	"hogo/core/responses"
)

func AttachInfoRoutes(app *components.Engine) {
	app.GET("/", responses.HandleAppInfo)
	app.GET("/info", responses.HandleAppInfo)
}
