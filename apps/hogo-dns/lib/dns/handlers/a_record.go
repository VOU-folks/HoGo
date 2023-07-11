package handlers

import (
	"net"

	"github.com/miekg/dns"
)

func HandleA(msg *dns.Msg, name string) {
	msg.Authoritative = true

	address := "127.0.0.1"
	msg.Answer = append(msg.Answer,
		&dns.A{
			Hdr: dns.RR_Header{Name: name, Rrtype: dns.TypeA, Class: dns.ClassINET, Ttl: 60},
			A:   net.ParseIP(address),
		},
	)
}
