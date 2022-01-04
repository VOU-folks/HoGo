package core

import (
  "fmt"
  "log"
  "os"
  "os/signal"
  "runtime"
  "syscall"
  "time"

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

func SetupCloseHandler() {
  c := make(chan os.Signal, 2)
  signal.Notify(c, os.Interrupt, syscall.SIGTERM)
  go func() {
    <-c
    fmt.Println("\rGot Interrupt Signal")

    time.Sleep(time.Second)
    defer components.GetServiceBus().Close()
    defer components.GetServiceBus().Drain()

    os.Exit(0)
  }()
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
        defer func() {
          for {
            time.Sleep(time.Millisecond * 300)
            if !components.GetServiceBus().Connected() {
              os.Exit(0)
            }
          }
        }()

        fmt.Println("\rGracefully closing connection to Service Bus")
        defer components.GetServiceBus().Close()
        defer components.GetServiceBus().Drain()
      }
    }
  }()

  runtime.Goexit()
}
