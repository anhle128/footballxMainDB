package models

import (
	"database/sql"
)

// IDatastorer declare base function for datastore
type IDatastorer interface {
	Gets(query string, args ...interface{}) (*sql.Rows, error)
	Get(query string, args ...interface{}) *sql.Row
	Insert(query string, args ...interface{}) (int64, error)
	Exec(query string, args ...interface{}) (int64, error)
}
