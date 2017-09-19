package models

// ICURD declare base function for each table
type ICURD interface {
	Create(model IQueryGenerator) (*IQueryGenerator, error)
	Update(modeDB IQueryGenerator) (int64, error)
	Delete(modeDB IQueryGenerator) (int64, error)
	Undo(modeDB IQueryGenerator) (int64, error)
	RealDelete(modeDB IQueryGenerator) (int64, error)
}
