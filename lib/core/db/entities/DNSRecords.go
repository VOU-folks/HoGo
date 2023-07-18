package entities

import "hogo/lib/core/db/enums"

type DNSRecords struct {
	Entity

	Id      string      `json:"id"`
	Zone    DNSZone     `json:"zone"`
	Records []DNSRecord `json:"records"`
	Active  bool        `json:"active"`
	Deleted bool        `json:"deleted"`
}

type DNSZone string

type DNSRecord struct {
	Type     enums.DNSRecordType `json:"type"`
	Name     string              `json:"name"`
	Value    string              `json:"value"`
	Priority int                 `json:"priority"`
}
