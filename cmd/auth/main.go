package main

import (
	"hogo/core"
	"hogo/services/auth"
)

func main() {
	core.LoadEnv()

	instance := auth.CreateInstance()
	auth.Start(instance)

	core.Wait(true)
}
