package http

import (
	. "hogo/lib/core/helpers"
)

type HTTPServer struct {
	stopped       bool
	stopRequested bool
}

func (s *HTTPServer) Init(args Args) {
	s.stopped = false
	s.stopRequested = false
}

func (s *HTTPServer) Listen() {
	// ToDo: Create db connections

	// ToDo: Create http listener

	// ToDo: Start http server listener
}

func (s *HTTPServer) Stop() {
	s.stopRequested = true

	// ToDo: Graceful shutdown logic

	s.stopped = true
}

func (s *HTTPServer) IsStopped() bool {
	return s.stopped
}
