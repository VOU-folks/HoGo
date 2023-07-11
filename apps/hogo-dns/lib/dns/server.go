package dns

import (
	"log"
	"strconv"
	"sync"

	"github.com/miekg/dns"
)

type Server struct {
	dnsService *dns.Server
	mx         sync.Mutex
}

func CreateServer() *Server {
	return &Server{
		dnsService: &dns.Server{
			Addr: "0.0.0.0:" + strconv.Itoa(53),
			Net:  "udp",
		},
		mx: sync.Mutex{},
	}
}

func (s *Server) SetHandler(handler dns.Handler) {
	s.dnsService.Handler = handler
}

func (s *Server) ListenAndServe() {
	log.Println("HoGo DNS Component listening at", s.dnsService.Addr)
	err := s.dnsService.ListenAndServe()
	if err != nil {
		log.Fatal("dnsService.ListenAndServe:", err.Error())
	}
}

func (s *Server) Close() {
	err := s.dnsService.Listener.Close()
	if err != nil {
		log.Fatal("dnsService.Listener:", err.Error())
	}
}
