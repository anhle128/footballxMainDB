package matches

import (
	"database/sql"
	"time"

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
func Create(model models.Match) (*models.Match, error) {
	newID, err := currentDS.Insert(model.GenCreateQuery())
	if err != nil {
		return nil, err
	}
	model.ID = newID
	return &model, nil
}

// Update club
// Implement ICURD interface
func Update(model models.Match) (int64, error) {
	return currentDS.Exec(model.GenUpdateQuery())
}

// Delete club
// Implement ICURD interface
func Delete(model models.Match) (int64, error) {
	return currentDS.Exec(model.GenDeleteQuery())
}

// Undo club
// Implement ICURD interface
func Undo(model models.Match) (int64, error) {
	return currentDS.Exec(model.GenUndoQuery())
}

// RealDelete club
// Implement ICURD interface
func RealDelete(model models.Match) (int64, error) {
	return currentDS.Exec(model.GenRealDeleteQuery())
}

//
// ─── SELECT FUNCTION ────────────────────────────────────────────────────────────
//

// GetBySeasenDateIndex get data by seasonID,  date, index
// return match by date and index
func GetBySeasenDateIndex(seasonID int64, date time.Time, index int64) (*models.Match, error) {
	row := currentDS.Get("Select * from matchs where index = $1 and date_unix = $2 and season_id = $3 and deleted = false", index, date.Unix(), seasonID)
	return convertRowToMatch(row)
}

// ForceGetBySeasenDateIndex force get data by easonID int, date time.Time, index int
// If not exist, create one and return
func ForceGetBySeasenDateIndex(seasonID int64, date time.Time, index int64) (*models.Match, error) {
	match, err := GetBySeasenDateIndex(seasonID, date, index)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return Create(models.Match{SeasonID: seasonID, Date: date, Index: index})
	}

	return match, nil
}

//
// ─── SUPPORT FUNCTION ───────────────────────────────────────────────────────────
//

func convertRowToMatch(row tables.IRowScanner) (*models.Match, error) {

	match := models.Match{}
	var strURLResult sql.NullString

	err := row.Scan(&match.ID,
		&match.SeasonID,
		&match.Date,
		&strURLResult,
		&match.Index,
		&match.DateUnix,
		&match.Deleted)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	if strURLResult.Valid {
		match.URLResult = strURLResult.String
	}

	return &match, nil
}
