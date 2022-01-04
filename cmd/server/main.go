package main

import (
  "log"
  "net/http"
  "os"

  "hogo/app/router"
  "hogo/core"
  "hogo/core/components"
  "hogo/core/helpers"
  "hogo/core/responses"

  "github.com/gin-gonic/gin"
)

func CreateApp() *components.Engine {
  app := components.NewHttpAppInstance()

  app.Use(
    gin.Logger(),
    responses.RecoverOnSystemError,
  )

  router.Attach(app)

  serviceBus := core.CreateServiceBus(os.Getenv("NATS_URL"))

  di := components.GetDI()
  di.ServiceBus = serviceBus
  di.App = app

  return app
}

func BindAppToSocket(app *components.Engine, socketAddr helpers.SocketAddr) {
  log.Println("Listening at:", socketAddr.ADDRESS)
  err := http.ListenAndServe(socketAddr.ADDRESS, app)
  if err != nil {
    panic(err)
  }
}

func main() {
  core.LoadEnv()
  go BindAppToSocket(CreateApp(), helpers.GetSocketAddr())
  core.Wait(true)
}
