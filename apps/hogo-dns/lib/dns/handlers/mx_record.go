package handlers

import (
	"github.com/miekg/dns"
)

func HandleMX(msg *dns.Msg, name string) {
	msg.Authoritative = true

	address := "127.0.0.1"
	msg.Answer = append(msg.Answer,
		&dns.MX{
			Hdr:        dns.RR_Header{Name: name, Rrtype: dns.TypeMX, Class: dns.ClassINET, Ttl: 60},
			Preference: 10,
			Mx:         address,
		},
	)
}
