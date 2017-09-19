package footballxMainDB

import (
	"database/sql"

	"github.com/anhle/footballxMainDB/models"
	_ "github.com/lib/pq"
)

type Datastore struct {
	*sql.DB
	*models.DBInfo
}

var currentDS Datastore

// var dbInfo *models.DBInfo

// InitTestDBInfo create db info for testing
func initTestDBInfo() *models.DBInfo {
	return &models.DBInfo{Username: "root", Password: "123456789", DB: "football-x-dev", Host: "localhost", Port: 5432}
}

// SetDBInfo need to opend connection
func SetDBInfo(info models.DBInfo) {
	currentDS.DBInfo = &info
}

// GetDBInfo return current db info
func GetDBInfo() models.DBInfo {
	return *currentDS.DBInfo
}

// Opend connection to db
func Opend() (*Datastore, error) {

	if currentDS.DBInfo == nil {
		currentDS.DBInfo = initTestDBInfo()
	}

	db, err := sql.Open("postgres", currentDS.DBInfo.GetConnectString())
	currentDS.DB = db

	if err != nil {
		return nil, err
	}

	return &currentDS, nil
}

// Gets multiple rows data
// Implement IDatastorer interface
func (ds Datastore) Gets(query string, args ...interface{}) (*sql.Rows, error) {
	return ds.DB.Query(query, args...)
}

// Get one row data
// Implement IDatastorer interface
func (ds Datastore) Get(query string, args ...interface{}) *sql.Row {
	return ds.DB.QueryRow(query, args...)
}

// Insert one row to database
// Implement IDatastorer interface
func (ds Datastore) Insert(query string, args ...interface{}) (int64, error) {

	var lastInsertID int64

	err := ds.DB.QueryRow(query, args...).Scan(&lastInsertID)
	if err != nil {
		return -1, err
	}
	return lastInsertID, nil
}

// Exec executes a prepared statement with the given arguments and
// returns a Result summarizing the effect of the statement.
// use fof update and delete query
// Implement IDatastorer interface
func (ds Datastore) Exec(query string, args ...interface{}) (int64, error) {
	stm, err := ds.DB.Prepare(query)
	if err != nil {
		return -1, err
	}

	result, err := stm.Exec(args...)
	if err != nil {
		return -1, err
	}

	return result.RowsAffected()
}

// // Create insert new data row to db
// // return id , error
// // Implement IDatastorer interface
// func (ds Datastore) Create(model models.IQueryGenerator) (int64, error) {
// 	return ds.Insert(model.GenCreateQuery())
// }

// // Update data to database
// // return number effected row, error
// // Implement IDatastorer interface
// func (ds Datastore) Update(modeDB models.IQueryGenerator) (int64, error) {
// 	return ds.Exec(modeDB.GenUpdateQuery())
// }

// // Delete data, change deleted = true
// // return number effected row, error
// // Implement IDatastorer interface
// func (ds Datastore) Delete(modeDB models.IQueryGenerator) (int64, error) {
// 	return ds.Exec(modeDB.GenDeleteQuery())
// }

// // Undo data, change deleted = false
// // return number effected row, error
// // Implement IDatastorer interface
// func (ds Datastore) Undo(modeDB models.IQueryGenerator) (int64, error) {
// 	return ds.Exec(modeDB.GenUndoQuery())
// }

// // ForceDelete data, real delete data from db
// // return number effected row, error
// // Implement IDatastorer interface
// func (ds Datastore) ForceDelete(modeDB models.IQueryGenerator) (int64, error) {
// 	return ds.Exec(modeDB.GenForceDeleteQuery())
// }

// Close connection to db
func Close() {
	if currentDS.DB != nil {
		currentDS.DB.Close()
	}
}
