package tables

// IRowScanner interface use for sql.Row and sql.Rows
type IRowScanner interface {
	Scan(dest ...interface{}) error
}
