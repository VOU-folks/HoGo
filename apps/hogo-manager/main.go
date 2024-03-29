package main

import (
	"embed"
	. "hogo/apps/hogo-manager/lib"
	"hogo/lib/core/globals"
	"hogo/lib/core/helpers"
	"hogo/lib/core/interfaces"
	"hogo/lib/core/log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var server interfaces.Server

func init() {
	InitGlobals()

	args := globals.Args
	log.Init(log.Level(args.Verbosity), log.Formatter(args.LogFormat))

	server = &HogoManagerServer{}
	server.Init(args)
}

func main() {
	go server.Listen()

	HandleInterruption()
}

func HandleInterruption() {
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

	sig := <-c
	log.Warn("Got ", sig, " signal")

	go server.Stop()

	WaitForShutdown()
}

func WaitForShutdown() {
	for {
		if server.IsStopped() {
			break
		}

		time.Sleep(time.Microsecond * 100)
	}

	log.Println("Bye (:")

	os.Exit(0)
}

/************/
/* GLOBALS */
/***********/

//go:embed app-info.json
var embeddedFiles embed.FS

func InitGlobals() {
	globals.Init(globals.InitArgs{
		EmbeddedFiles: embeddedFiles,
	})

	helpers.PrintAppInfo(globals.AppInfo)
	helpers.PrintArgs(globals.Args)
}
