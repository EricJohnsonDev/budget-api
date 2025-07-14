package models

import "database/sql"

type Tx_expenses struct {
	ID          int
	Date        string
	Amount      string
	Institution string
	Category    string
	Subcategory sql.NullString
	Comment     string
}
