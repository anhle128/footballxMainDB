package clubs

import (
	"testing"

	database "github.com/anhle/footballxMainDB"
)

func TestGetByID(t *testing.T) {
	database.Opend()
	defer database.Close()

	_, err := GetByID(1)
	if err != nil {
		t.Error(err)
	}
}

func TestGets(t *testing.T) {
	database.Opend()
	defer database.Close()

	_, err := Gets("select * from clubs")
	if err != nil {
		t.Error(err)
	}
}
