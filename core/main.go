package core

import (
  "fmt"
  "log"
  "os"
  "os/signal"
  "runtime"
  "syscall"

  "hogo/core/components"

  "github.com/joho/godotenv"
)

func LoadEnv() {
  err := godotenv.Load()
  if err != nil {
    fmt.Println(err.Error())
  }
}

func CreateServiceBus(serviceBusAddress string) *components.ServiceBus {
  serviceBus := components.GetServiceBus()
  serviceBus.Connect(serviceBusAddress)

  log.Println("Connected to Service Bus")
  return serviceBus
}

func Shutdown(graceful bool) {
  if graceful {
    if components.GetServiceBus().Connected() {
      fmt.Println("\rDraining connection to Service Bus")
      components.GetServiceBus().Drain()

      fmt.Println("\rClosing connection to Service Bus")
      components.GetServiceBus().Close()
    }
  }
  os.Exit(0)
}

func Wait(gracefulShutdownOnExit bool) {
  log.Println("Press Control-C to stop")

  c := make(chan os.Signal, 2)
  signal.Notify(
    c,
    syscall.SIGINT,
    syscall.SIGTERM,
    syscall.SIGKILL,
    syscall.SIGUSR1,
    syscall.SIGUSR2,
  )
  go func() {
    sig := <-c
    fmt.Println("\rGot", sig, "signal")

    Shutdown(gracefulShutdownOnExit)
  }()

  runtime.Goexit()
}
