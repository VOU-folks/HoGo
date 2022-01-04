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

func Wait(gracefulShutdownOnExit bool) {
  log.Println("Press Control-C to stop")

  c := make(chan os.Signal, 2)
  signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGUSR1, syscall.SIGUSR2)
  go func() {
    <-c
    fmt.Println("\rGot interrupt signal")

    if gracefulShutdownOnExit {
      if components.GetServiceBus().Connected() {
        fmt.Println("\rDraining connection to Service Bus")
        components.GetServiceBus().Drain()

        fmt.Println("\rClosing connection to Service Bus")
        components.GetServiceBus().Close()
      }
    }

    os.Exit(0)
  }()

  runtime.Goexit()
}
