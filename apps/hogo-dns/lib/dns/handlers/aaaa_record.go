package handlers

import (
	"net"

	"github.com/miekg/dns"
)

func HandleAAAA(msg *dns.Msg, name string) {
	msg.Authoritative = true

	address := "::ffff:7f00:1"
	msg.Answer = append(msg.Answer,
		&dns.AAAA{
			Hdr:  dns.RR_Header{Name: name, Rrtype: dns.TypeAAAA, Class: dns.ClassINET, Ttl: 60},
			AAAA: net.ParseIP(address),
		},
	)
}
