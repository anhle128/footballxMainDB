package models

type MatchClub struct {
	club_id int
	match_id int
	home bool
	goals int
	result int
	deleted bool
}