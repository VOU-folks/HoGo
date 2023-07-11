package main

import (
	. "hogo/apps/hogo-manager/lib"
	"hogo/lib/core/helpers"
	"hogo/lib/core/log"
)

func main() {
	args := helpers.GetArgs()
	log.Init(log.Level(args.Verbosity), log.Formatter(args.LogFormat))

	server := &HogoManagerServer{}
	server.Init(args)
	server.Listen()
}
