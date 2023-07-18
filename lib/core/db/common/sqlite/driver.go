package sqlite

import (
	"database/sql"
	"errors"
	_ "github.com/glebarez/go-sqlite"
	"hogo/lib/core/helpers"
	"hogo/lib/core/log"
)

var DB *sql.DB

type SQLiteConnector struct {
}

func (d SQLiteConnector) Open() error {
	log.Debug("Opening database:", helpers.GetEnv("CORE_DB_PATH"))

	db, err := sql.Open("sqlite", helpers.GetEnv("CORE_DB_PATH")+"?_pragma=journal_mode(OFF)")
	if err != nil {
		return err
	}

	DB = db

	return nil
}

func (d SQLiteConnector) Close() error {
	if DB != nil {
		return nil
	}
	return DB.Close()
}

func (d SQLiteConnector) GetDB() (*sql.DB, error) {
	if DB == nil {
		return nil, errors.New("DB connection not ready")
	}

	return DB, nil
}
