package models

// MatchClub model
type MatchClub struct {
	ID      int64
	ClubID  int64
	MatchID int64
	Home    bool
	Goals   int64
	Result  int64
	Deleted bool
}
