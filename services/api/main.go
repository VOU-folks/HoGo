package api

import (
  "log"
  "net/http"
  "os"

  "github.com/gin-gonic/gin"
  "hogo/core"
  "hogo/core/components"
  "hogo/core/helpers"
  "hogo/core/responses"
  "hogo/services/api/router"
)

func CreateInstance() *components.DI {
  di := components.GetDI()

  di.ServiceBus = core.CreateServiceBus(os.Getenv("NATS_URL"))

  di.App = components.NewHttpAppInstance()
  di.App.Use(
    gin.Logger(),
    responses.RecoverOnSystemError,
  )

  return di
}

func BindAppToSocket(app *components.Engine, socketAddr helpers.SocketAddr) {
  log.Println("Listening at:", socketAddr.ADDRESS)
  err := http.ListenAndServe(socketAddr.ADDRESS, app)
  if err != nil {
    panic(err)
  }
}

func Start(app *components.Engine) {
  router.Attach(app)
}
