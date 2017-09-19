package footballxMainDB

import (
	"testing"
)

func TestOpendAndCloseDB(t *testing.T) {
	err := Opend()

	if err != nil {
		t.Error("Cannot opend connection with info:", GetDBInfo())
	}

	defer Close()
}

func TestGetDatas(t *testing.T) {
	errOpend := Opend()

	if errOpend != nil {
		t.Error("Cannot opend connection with info:", GetDBInfo())
	}

	defer Close()

	_, errQuery := currentDS.Gets("SELECT * FROM clubs;")
	checkErr(t, errQuery)
}

func TestCountData(t *testing.T) {
	errOpend := Opend()

	if errOpend != nil {
		t.Error("Cannot opend connection with info:", GetDBInfo())
	}

	defer Close()

	row := currentDS.Get("SELECT count(*) FROM clubs")

	var totalNumberClubs int
	row.Scan(&totalNumberClubs)
}

func Test_Insert_Update_Delete_Data(t *testing.T) {

	errOpend := Opend()

	if errOpend != nil {
		t.Error("Cannot opend connection with info:", GetDBInfo())
	}

	defer Close()

	newID, errInsert := currentDS.Insert("INSERT INTO clubs (name) VALUES ($1) returning id", "testinsert")
	checkErr(t, errInsert)

	_, errUpdate := currentDS.Exec("UPDATE clubs SET deleted = $2 WHERE id = $1", newID, true)
	checkErr(t, errUpdate)

	_, errDelete := currentDS.Exec("DELETE FROM clubs WHERE clubs.id = $1", newID)
	checkErr(t, errDelete)
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
		return
	}
}
