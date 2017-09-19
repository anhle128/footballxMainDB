package models

// IQueryGenerator interface
// Declare base function generate query for model
type IQueryGenerator interface {
	GenCreateQuery() string
	GenUpdateQuery() string
	GenDeleteQuery() string
	GenRealDeleteQuery() string
	GenUndoQuery() string
}
