package lib

import (
	"context"
	"errors"
	"hogo/apps/hogo-manager/lib/counters"
	. "hogo/lib/core/helpers"
	. "hogo/lib/core/http"
	"hogo/lib/core/http/handlers"
	"hogo/lib/core/interfaces"
	"hogo/lib/core/log"
	"net/http"
	"runtime"
	"time"
)

type HogoManagerServer struct {
	stopped       bool
	stopRequested bool

	app        interfaces.App
	httpServer *http.Server
}

func (s *HogoManagerServer) Init(args Args) {
	s.stopped = false
	s.stopRequested = false

	s.httpServer = CreateHttpServer(args.ListenAt)
	s.app = &HogoManagerApp{}
	s.app.Init(args)
	s.app.BindToHttpServer(s.httpServer)
}

func (s *HogoManagerServer) Listen() {
	if s.stopRequested {
		log.Error("HogoManagerServer.Listen:", errors.New("stop requested before listen call"))
	}

	err := s.httpServer.ListenAndServe()
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("httpServer.ListenAndServe:", err.Error())
		}
	}
}

func (s *HogoManagerServer) Stop() {
	if s.stopRequested {
		return
	}
	s.stopRequested = true

	for {
		if counters.ActiveRequests.Value() == 0 {
			break
		}

		time.Sleep(time.Microsecond * 100)
	}

	s.httpServer.Handler = &handlers.ShutdownHandler{}
	time.Sleep(time.Second * 5)
	err := s.httpServer.Shutdown(context.Background())
	if err != nil {
		log.Fatal("httpServer.Close:", err.Error())
	}

	s.stopped = true

	runtime.Goexit()
}

func (s *HogoManagerServer) IsStopped() bool {
	return s.stopped
}
