package models

import "time"

type Match struct {
	id int
	season_id int
	time string
	date time.Time
	deleted bool
}
