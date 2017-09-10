package clubs

import (
	"database/sql"

	database "github.com/anhle/footballxMainDB"
	"github.com/anhle/footballxMainDB/models"
	"github.com/anhle/footballxMainDB/tables"
)

// GetByID return data by id
func GetByID(id int) (*models.Club, error) {
	row := database.Get("Select * from clubs where id = $1", id)
	return convertRowToClubs(row)
}

// Gets return mitiple datas
func Gets(query string, args ...interface{}) ([]models.Club, error) {
	rows, err := database.Gets(query, args...)
	if err != nil {
		return nil, err
	}

	var clubs []models.Club

	for rows.Next() {
		club, err := convertRowToClubs(rows)
		if err != nil {
			return nil, err
		}
		clubs = append(clubs, *club)
	}

	return clubs, nil
}

func convertRowToClubs(row tables.IRowScan) (*models.Club, error) {

	club := models.Club{}
	var strIcon sql.NullString

	err := row.Scan(&club.ID, &club.Name, &strIcon, &club.Deleted)
	if err != nil {
		return nil, err
	}

	if strIcon.Valid {
		club.Icon = strIcon.String
	}

	return &club, nil
}
