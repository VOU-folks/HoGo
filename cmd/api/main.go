package main

import (
	"hogo/core"
	"hogo/core/helpers"
	"hogo/services/api"
)

func main() {
	core.LoadEnv()

	instance := api.CreateInstance()
	go api.BindAppToSocket(instance.App, helpers.GetSocketAddr())
	api.Start(instance)

	core.Wait(true)
}
