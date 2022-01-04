package router

import (
  "hogo/core/components"
  "hogo/core/responses"
)

func AttachNotFound(app *components.Engine) {
  app.NoRoute(responses.HandleNotFound)
}
