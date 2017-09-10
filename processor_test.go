package footballxMainDB

import "testing"

func InitDBInfo() DatabaseInfo {
	return DatabaseInfo{Username: "root", Password: "123456789", DB: "football-x-dev", Host: "localhost", Port: 5432}
}

func TestOpendAndCloseDB(t *testing.T) {

	dbInfo := InitDBInfo()

	SetDatabaseInfo(dbInfo)

	err := Opend()

	if err != nil {
		t.Error("Cannot opend connection with info:", dbInfo)
	}

	defer Close()
}

func TestGetDatas(t *testing.T) {
	dbInfo := InitDBInfo()

	SetDatabaseInfo(dbInfo)

	errOpend := Opend()

	if errOpend != nil {
		t.Error("Cannot opend connection with info:", dbInfo)
	}

	defer Close()

	_, errQuery := Gets("SELECT * FROM clubs;")
	checkErr(t, errQuery)
}

func TestGetData(t *testing.T) {
	dbInfo := InitDBInfo()

	SetDatabaseInfo(dbInfo)

	errOpend := Opend()

	if errOpend != nil {
		t.Error("Cannot opend connection with info:", dbInfo)
	}

	defer Close()

	row := Get("SELECT count(*) FROM clubs")

	var totalNumberClubs int
	row.Scan(&totalNumberClubs)
}

func Test_Insert_Update_Delete_Data(t *testing.T) {
	dbInfo := InitDBInfo()

	SetDatabaseInfo(dbInfo)

	errOpend := Opend()

	if errOpend != nil {
		t.Error("Cannot opend connection with info:", dbInfo)
	}

	defer Close()

	newID, errInsert := Insert("INSERT INTO clubs (name) VALUES ($1) returning id", "testinsert")
	checkErr(t, errInsert)

	_, errUpdate := Delete("UPDATE clubs SET deleted = $2 WHERE id = $1", newID, true)
	checkErr(t, errUpdate)

	_, errDelete := Delete("DELETE FROM clubs WHERE clubs.id = $1", newID)
	checkErr(t, errDelete)
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		return
	}
}
