package models

import (
	"fmt"
)

// Club model
type Club struct {
	ID      int64
	Name    string
	Icon    string
	Deleted bool
}

// GenCreateQuery from data
// Implement IQueryGenerator interface
func (c Club) GenCreateQuery() string {
	if len(c.Icon) != 0 {
		return fmt.Sprintf("INSERT INTO clubs (name,icon) VALUES ('%s','%s')", c.Name, c.Icon)
	}
	return fmt.Sprintf("INSERT INTO clubs (name) VALUES ('%s')", c.Name)
}

// GenUpdateQuery from data
// Implement IQueryGenerator interface
func (c Club) GenUpdateQuery() string {
	return fmt.Sprintf("UPDATE clubs SET name = '%s', icon = '%s' WHERE id = %d", c.Name, c.Icon, c.ID)
}

// GenDeleteQuery set deleted = true
// Implement IQueryGenerator interface
func (c Club) GenDeleteQuery() string {
	return fmt.Sprintf("UPDATE clubs SET deleted = true WHERE id = %d", c.ID)
}

// GenUndoQuery set deleted = false
// Implement IQueryGenerator interface
func (c Club) GenUndoQuery() string {
	return fmt.Sprintf("UPDATE clubs SET deleted = false WHERE id = %d", c.ID)
}

// GenRealDeleteQuery real delete
// Implement IQueryGenerator interface
func (c Club) GenRealDeleteQuery() string {
	return fmt.Sprintf("DELETE FROM clubs WHERE clubs.id = %d", c.ID)
}
