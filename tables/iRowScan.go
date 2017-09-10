package tables

// IRowScan interface use for sql.Row and sql.Rows
type IRowScan interface {
	Scan(dest ...interface{}) error
}
