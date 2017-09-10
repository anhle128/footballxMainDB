package models

import "time"

type season struct{
	int int
	league_id string
	date_start time.Time
	date_end time.Time
	deleted bool
}

