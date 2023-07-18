package sqlite

import (
	"context"
	"github.com/doug-martin/goqu/v9"
	. "hogo/lib/core/db/common/dao"
	"hogo/lib/core/db/entities"
	"strconv"

	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	"github.com/georgysavva/scany/sqlscan"
)

type SQLiteDao[T entities.IEntity] struct {
	Dao[T]

	collectionName string
}

func build() goqu.DialectWrapper {
	return goqu.Dialect("sqlite3")
}

func (d *SQLiteDao[T]) SetCollectionName(name string) {
	d.collectionName = name
}

func (d *SQLiteDao[T]) CollectionName() string {
	return d.collectionName
}

func (d *SQLiteDao[T]) FindById(id string) (*T, error) {
	var dest []*T

	query, _, _ := build().From(d.CollectionName()).Where(goqu.Ex{"id": id}).Limit(1).ToSQL()
	err := sqlscan.Select(context.Background(), DB, &dest, query)
	if err != nil {
		return nil, err
	}

	if len(dest) == 1 {
		return dest[0], nil
	}
	return nil, nil
}

func (d *SQLiteDao[T]) Create(data interface{}) (*T, error) {
	query, _, _ := build().Insert(d.CollectionName()).Rows(data).ToSQL()
	result, err := DB.Exec(query)
	if err != nil {
		return nil, err
	}

	id, _ := result.LastInsertId()
	return d.FindById(strconv.FormatInt(id, 10))
}
