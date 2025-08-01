package models

type Tx_expenses struct {
	ID          int
	Date        string
	Amount      string
	Institution string
	Category    string
	// *string in place of sql.NullString
	// See https://stackoverflow.com/questions/40092155/difference-between-string-and-sql-nullstring
	// *string seems much easier to work with at this point
	Subcategory *string
	Comment     string
}
