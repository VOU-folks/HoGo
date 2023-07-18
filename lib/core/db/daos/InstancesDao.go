package daos

import (
	"hogo/lib/core/db/common/sqlite"
	"hogo/lib/core/db/entities"
)

type InstancesDao struct {
	sqlite.SQLiteDao[entities.Instance]
}

func NewInstancesDao() *InstancesDao {
	instance := &InstancesDao{}
	instance.SetCollectionName("instances")

	return instance
}
