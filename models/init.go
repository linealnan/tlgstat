package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB

func Init(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
