package dao

import (
	"hogo/lib/core/db/entities"
)

type Dao[T entities.IEntity] interface {
	SetCollectionName(name string)
	CollectionName() string

	Create(data interface{}) (*T, error)
	FindById(id string) (*T, error)
}
