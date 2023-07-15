package main

import (
	"embed"
	. "hogo/apps/hogo-manager/lib"
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

	log.Init(log.Level(Args.Verbosity), log.Formatter(Args.LogFormat))

	server = &HogoManagerServer{}
	server.Init(Args)
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

var AppInfo helpers.AppInfo
var Args helpers.Args

//go:embed app-info.json
var embeddedFiles embed.FS

func InitGlobals() {
	AppInfo = helpers.GetAppInfo(embeddedFiles)
	helpers.PrintAppInfo(AppInfo)

	Args = helpers.GetArgs()
	helpers.PrintArgs(Args)
}
