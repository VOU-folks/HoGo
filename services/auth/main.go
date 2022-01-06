package auth

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"

	"hogo/core"
	"hogo/core/components"
)

func CreateInstance() *components.DI {
	di := components.GetDI()

	serviceBus := core.CreateServiceBus(os.Getenv("NATS_URL"))
	di.ServiceBus = serviceBus

	return di
}

func Start(instance *components.DI) {
	instance.ServiceBus.Subscribe("auth.attempt", func(msg *nats.Msg) {
		fmt.Println(msg)
	})
}
