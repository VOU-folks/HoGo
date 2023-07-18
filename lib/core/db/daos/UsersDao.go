package daos

import (
	"hogo/lib/core/db/common/sqlite"
	"hogo/lib/core/db/entities"
)

type UsersDao struct {
	sqlite.SQLiteDao[entities.User]
}

func NewUsersDao() *UsersDao {
	instance := &UsersDao{}
	instance.SetCollectionName("users")

	return instance
}
