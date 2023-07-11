package dns

import (
	"log"

	"github.com/miekg/dns"

	"hogo/apps/hogo-dns/lib/dns/handlers"
)

type Handler struct {
}

func (h Handler) ServeDNS(res dns.ResponseWriter, req *dns.Msg) {
	msg := handlers.HandleDNSQuery(req)
	err := res.WriteMsg(msg)
	if err != nil {
		log.Print("Got error while sending DNS response: ", err.Error())
	}
}
