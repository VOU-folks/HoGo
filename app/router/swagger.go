package router

import (
  swaggerFiles "github.com/swaggo/files"
  ginSwagger "github.com/swaggo/gin-swagger"

  "hogo/core/components"
)

func AttachSwaggerRoutes(app *components.Engine) {
  swaggerUrl := ginSwagger.URL("/docs/swagger/doc.json")
  app.GET("/docs/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, swaggerUrl))
}
