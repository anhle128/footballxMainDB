package footballxMainDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DatabaseInfo data need to connect db
type DatabaseInfo struct {
	Username string
	Password string
	DB       string
	Host     string
	Port     int
}

var currentDB *sql.DB
var databaseInfo DatabaseInfo

func (info DatabaseInfo) getConnectString() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable host=%s port=%d",
		info.Username, info.Password, info.DB, info.Host, info.Port)
}

// SetDatabaseInfo need to opend connection
func SetDatabaseInfo(info DatabaseInfo) {
	databaseInfo = info
}

// Opend connection to db
func Opend() error {

	db, err := sql.Open("postgres", databaseInfo.getConnectString())
	currentDB = db

	if err != nil {
		return err
	}

	return nil
}

// Gets return multiple rows
func Gets(query string, args ...interface{}) (*sql.Rows, error) {
	return currentDB.Query(query, args...)
}

// Get return one row
func Get(query string, args ...interface{}) *sql.Row {
	return currentDB.QueryRow(query, args...)
}

// Insert one row to database
func Insert(query string, args ...interface{}) (int, error) {

	var lastInsertID int

	err := currentDB.QueryRow(query, args...).Scan(&lastInsertID)
	if err != nil {
		return -1, err
	}
	return lastInsertID, nil
}

// Delete one or more rows data
func Delete(query string, args ...interface{}) (int64, error) {
	return execDatas(query, args...)
}

// Update one or more rows data
func Update(query string, args ...interface{}) (int64, error) {
	return execDatas(query, args...)
}

func execDatas(query string, args ...interface{}) (int64, error) {
	stm, err := currentDB.Prepare(query)
	if err != nil {
		return -1, err
	}

	result, err := stm.Exec(args...)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

// Close connection to db
func Close() {
	if currentDB != nil {
		currentDB.Close()
		currentDB = nil
	}
}
