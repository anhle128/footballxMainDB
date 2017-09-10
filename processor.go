package footballxMainDB

import (
	"database/sql"

	"github.com/anhle/footballxMainDB/models"
	_ "github.com/lib/pq"
)

var currentDB *sql.DB
var dbInfo *models.DBInfo

// InitTestDBInfo create db info for testing
func initTestDBInfo() *models.DBInfo {
	return &models.DBInfo{Username: "root", Password: "123456789", DB: "football-x-dev", Host: "localhost", Port: 5432}
}

// SetDBInfo need to opend connection
func SetDBInfo(info models.DBInfo) {
	dbInfo = &info
}

// GetDBInfo return current db info
func GetDBInfo() models.DBInfo {
	return *dbInfo
}

// Opend connection to db
func Opend() error {

	if dbInfo == nil {
		dbInfo = initTestDBInfo()
	}

	db, err := sql.Open("postgres", dbInfo.GetConnectString())
	currentDB = db

	if err != nil {
		return err
	}

	return nil
}

// Gets multiple rows data
func Gets(query string, args ...interface{}) (*sql.Rows, error) {
	return currentDB.Query(query, args...)
}

// Get one row data
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

// Exec executes a prepared statement with the given arguments and
// returns a Result summarizing the effect of the statement.
// use fof update and delete query
func Exec(query string, args ...interface{}) (int64, error) {
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
