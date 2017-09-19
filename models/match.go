package models

import (
	"fmt"
	"time"
)

// Match model
type Match struct {
	ID        int64
	SeasonID  int64
	Date      time.Time
	DateUnix  int64
	URLResult string
	Index     int64
	Deleted   bool
}

// GenCreateQuery from data
// Implement QueryGenerator  interface
func (m Match) GenCreateQuery() string {
	return fmt.Sprintf("INSERT INTO matchs (season_id,date,url_result,index,date_unix) VALUES (%d,'%s','%s',%d,%d)",
		m.SeasonID, m.Date.Format("2006-01-02 15:04:00"), m.URLResult, m.Index, m.Date.Unix())
}

// GenUpdateQuery from data
// Implement QueryGenerator  interface
func (m Match) GenUpdateQuery() string {
	return fmt.Sprintf("UPDATE matchs SET season_id = %d, date='%s', url_result ='%s', date_unix = %d WHERE id = %d",
		m.SeasonID, m.Date.Format("2006-01-02 15:04:00"), m.URLResult, m.Date.Unix(), m.ID)
}

// GenDeleteQuery set deleted = true
// Implement QueryGenerator  interface
func (m Match) GenDeleteQuery() string {
	return fmt.Sprintf("UPDATE matchs SET deleted = true WHERE id = %d", m.ID)
}

// GenUndoQuery set deleted = false
// Implement QueryGenerator  interface
func (m Match) GenUndoQuery() string {
	return fmt.Sprintf("UPDATE matchs SET deleted = false WHERE id = %d", m.ID)
}

// GenRealDeleteQuery real delete
// Implement QueryGenerator  interface
func (m Match) GenRealDeleteQuery() string {
	return fmt.Sprintf("DELETE FROM matchs WHERE clubs.id = %d", m.ID)
}
