package enums

type DNSRecordType string

const (
	DNSRecordTypeA     = "A"
	DNSRecordTypeAAAA  = "AAAA"
	DNSRecordTypeCNAME = "CNAME"
	DNSRecordTypeTXT   = "TXT"
	DNSRecordTypeMX    = "MX"
	DNSRecordTypePTR   = "PTR"
)
