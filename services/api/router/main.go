package router

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"hogo/core/components"
	"hogo/core/responses"

	v1 "hogo/services/api/router/v1"
)

func Attach(app *components.Engine) {
	app.NoRoute(responses.HandleNotFound)

	var router = app.Group("/")
	router.GET("/", responses.HandleAppInfo)
	router.GET("/info", responses.HandleAppInfo)
	router.GET("/docs/swagger/*any",
		ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/docs/swagger/doc.json")),
	)

	var v1Router = router.Group("/v1")
	v1.AttachV1Routes(app, v1Router)

}
