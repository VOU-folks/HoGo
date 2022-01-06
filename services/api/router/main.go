package router

import (
	"hogo/core/components"
)

func Attach(app *components.Engine) {
	AttachInfoRoutes(app)
	AttachSwaggerRoutes(app)
	AttachNotFound(app)
}
