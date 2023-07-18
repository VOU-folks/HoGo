package daos

import (
	"hogo/lib/core/db/common/sqlite"
	"hogo/lib/core/db/entities"
)

type DNSRecordsDao struct {
	sqlite.SQLiteDao[entities.DNSRecord]
}

func NewDNSRecordsDao() *DNSRecordsDao {
	instance := &DNSRecordsDao{}
	instance.SetCollectionName("dns_records")

	return instance
}
