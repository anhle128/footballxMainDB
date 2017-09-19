package models

import "time"

// Season model
type Season struct {
	ID        int64
	LeagueID  int64
	DateStart time.Time
	DateEnd   time.Time
	Deleted   bool
}
