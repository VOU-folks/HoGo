package daos

import (
	"hogo/lib/core/db/common/sqlite"
	"hogo/lib/core/db/entities"
)

type PlacementsDao struct {
	sqlite.SQLiteDao[entities.Placement]
}

func NewPlacementsDao() *PlacementsDao {
	instance := &PlacementsDao{}
	instance.SetCollectionName("placements")

	return instance
}
