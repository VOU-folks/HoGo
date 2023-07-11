package lib

import (
	. "hogo/lib/core/helpers"
	. "hogo/lib/core/http"
	"hogo/lib/core/log"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

type HogoManagerServer struct {
	HTTPServer
	stopped       bool
	stopRequested bool
}

func (s *HogoManagerServer) Init(args Args) {
	s.stopped = false
	s.stopRequested = false

	s.HTTPServer.Init(args)

	// ToDo: Setup necessary preparations before Listen() call
}

func (s *HogoManagerServer) Listen() {
	s.HTTPServer.Listen()

	s.HandleInterruption()
}

func (s *HogoManagerServer) HandleInterruption() {
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

	go s.Shutdown()

	s.ShutdownHandler()
}

func (s *HogoManagerServer) ShutdownHandler() {
	for {
		if s.IsStopped() {
			break
		}

		time.Sleep(time.Microsecond * 100)
	}

	log.Println("Bye (:")

	os.Exit(0)
}

func (s *HogoManagerServer) Shutdown() {
	s.stopRequested = true

	// ToDo: Graceful shutdown logic
	/*
		Expected to have code of closing db connections, tasks and etc
	*/
	time.Sleep(time.Second)

	s.HTTPServer.Stop()

	runtime.Goexit()
}
