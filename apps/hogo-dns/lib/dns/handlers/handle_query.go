package handlers

import (
	"log"

	"github.com/miekg/dns"
)

func HandleDNSQuery(req *dns.Msg) *dns.Msg {
	question := req.Question[0]
	recordType := question.Qtype
	recordName := question.Name

	log.Println("Got DNS Request: ", recordType, recordName)

	msg := &dns.Msg{}
	msg.SetReply(req)

	switch recordType {
	case dns.TypeA:
		HandleA(msg, recordName)

	case dns.TypeAAAA:
		HandleAAAA(msg, recordName)

	case dns.TypeCNAME:
		HandleCNAME(msg, recordName)

	case dns.TypeMX:
		HandleMX(msg, recordName)
	}

	return msg
}
