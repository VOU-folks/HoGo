package daos

import (
	"hogo/lib/core/db/common/sqlite"
	"hogo/lib/core/db/entities"
)

type ProjectsDao struct {
	sqlite.SQLiteDao[entities.Project]
}

func NewProjectsDao() *ProjectsDao {
	instance := &ProjectsDao{}
	instance.SetCollectionName("projects")

	return instance
}
