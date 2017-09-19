package clubs

import (
	"database/sql"

	"github.com/anhle/footballxMainDB/models"
	"github.com/anhle/footballxMainDB/tables"
)

//
// ─── DEPENDENCE INJECTION ───────────────────────────────────────────────────────
//

var currentDS models.IDatastorer

// SetDatastore injection datastore
func SetDatastore(ds models.IDatastorer) {
	currentDS = ds
}

//
// ─── SIMPLE CURD FUNCTION ───────────────────────────────────────────────────────
//

// Create club
// Implement ICURD interface
func Create(model models.Club) (*models.Club, error) {
	newID, err := currentDS.Insert(model.GenCreateQuery())
	if err != nil {
		return nil, err
	}
	model.ID = newID
	return &model, nil
}

// Update club
// Implement ICURD interface
func Update(model models.Club) (int64, error) {
	return currentDS.Exec(model.GenUpdateQuery())
}

// Delete club
// Implement ICURD interface
func Delete(model models.Club) (int64, error) {
	return currentDS.Exec(model.GenDeleteQuery())
}

// Undo club
// Implement ICURD interface
func Undo(model models.Club) (int64, error) {
	return currentDS.Exec(model.GenUndoQuery())
}

// RealDelete club
// Implement ICURD interface
func RealDelete(model models.Club) (int64, error) {
	return currentDS.Exec(model.GenRealDeleteQuery())
}

//
// ─── SELECT FUNCTION ────────────────────────────────────────────────────────────
//

// GetByID get club by id
func GetByID(id int) (*models.Club, error) {
	row := currentDS.Get("Select * from clubs where id = $1", id)
	return convertRowToClub(row)
}

// GetByName get clubs by name
func GetByName(name string) (*models.Club, error) {
	row := currentDS.Get("Select * from clubs where name = $1 and deleted = false", name)
	return convertRowToClub(row)
}

// ForceGetByName force get club by name
// if is exsit return club
// if not, create and return
func ForceGetByName(name string) (*models.Club, error) {
	pointerClub, err := GetByName(name)

	if err != nil {
		return nil, err
	}

	if pointerClub != nil {
		return pointerClub, nil
	}

	return Create(models.Club{Name: name, Deleted: false})
}

//
// ─── SUPPORT FUNCTION ───────────────────────────────────────────────────────────
//

func convertRowToClub(row tables.IRowScanner) (*models.Club, error) {

	club := models.Club{}
	var strIcon sql.NullString

	err := row.Scan(&club.ID, &club.Name, &strIcon, &club.Deleted)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if strIcon.Valid {
		club.Icon = strIcon.String
	}

	return &club, nil
}
