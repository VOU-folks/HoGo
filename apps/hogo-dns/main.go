package main

import (
	"os"
	"os/signal"
	"runtime"
	"syscall"

	"hogo/apps/hogo-dns/lib/api"
	"hogo/apps/hogo-dns/lib/dns"
	"hogo/lib/core/log"
)

var dnsServer *dns.Server
var apiServer *api.Server

func main() {

	dnsServer = dns.CreateServer()
	dnsServer.SetHandler(dns.Handler{})

	apiServer = api.CreateServer()

	go dnsServer.ListenAndServe()
	go apiServer.ListenAndServe()

	WaitForCloseSignal(true)
}

func WaitForCloseSignal(gracefulShutdownOnExit bool) {
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
	log.WithFields(log.Fields{
		"sig": sig,
	}).Println("Got", sig, "signal")

	Shutdown()

	runtime.Goexit()
}

func Shutdown() {
	dnsServer.Close()
	apiServer.Close()
	
	log.Println("Bye (:")

	os.Exit(0)
}
