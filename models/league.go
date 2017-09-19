package models

// League model
type League struct {
	ID      int64
	Name    string
	Icon    string
	URL     string
	ArenaID string
	Deleted bool
}
