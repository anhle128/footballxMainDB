package clubs

import (
	"testing"

	database "github.com/anhle/footballxMainDB"
)

func TestGetByID(t *testing.T) {
	ds, _ := database.Opend()
	defer database.Close()
	SetDatastore(ds)

	_, err := GetByID(1)
	if err != nil {
		t.Error(err)
	}
}

func TestGetByName(t *testing.T) {
	ds, _ := database.Opend()
	defer database.Close()
	SetDatastore(ds)

	_, err := GetByName("Manchester United")
	if err != nil {
		t.Error(err)
	}
}

func TestCheckClub_Delete_ForceDelete(t *testing.T) {
	ds, _ := database.Opend()
	defer database.Close()
	SetDatastore(ds)

	club, err := ForceGetByName("TestCheckClub")
	if err != nil {
		t.Error(err)
	}

	_, err2 := Delete(*club)
	if err2 != nil {
		t.Error(err2)
	}

	_, err3 := RealDelete(*club)
	if err3 != nil {
		t.Error(err3)
	}

}

func TestGets(t *testing.T) {
	database.Opend()
	defer database.Close()

	_, err := currentDS.Gets("select * from clubs")
	if err != nil {
		t.Error(err)
	}
}
