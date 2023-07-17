package types

type DNSRecords struct {
	Id      string      `json:"id"`
	Zone    DNSZone     `json:"zone"`
	Records []DNSRecord `json:"records"`
	Active  bool        `json:"active"`
	Deleted bool        `json:"deleted"`
}

type DNSZone string

type DNSRecord struct {
	Type     DNSRecordType `json:"type"`
	Name     string        `json:"name"`
	Value    string        `json:"value"`
	Priority int           `json:"priority"`
}

type DNSRecordType string

const (
	DNSRecordTypeA     = "A"
	DNSRecordTypeAAAA  = "AAAA"
	DNSRecordTypeCNAME = "CNAME"
	DNSRecordTypeTXT   = "TXT"
	DNSRecordTypeMX    = "MX"
	DNSRecordTypePTR   = "PTR"
)
