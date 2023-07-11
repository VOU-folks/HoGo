package handlers

import (
	"github.com/miekg/dns"
)

func HandleCNAME(msg *dns.Msg, name string) {
	msg.Authoritative = true

	target := "localhost"
	msg.Answer = append(msg.Answer,
		&dns.CNAME{
			Hdr:    dns.RR_Header{Name: name, Rrtype: dns.TypeCNAME, Class: dns.ClassINET, Ttl: 60},
			Target: target,
		},
	)
}
