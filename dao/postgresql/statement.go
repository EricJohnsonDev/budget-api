package postgresql

const INSERT_EXPENSE = "INSERT INTO tx_expenses (\"Date\", \"Amount\", \"Institution\", \"Category\", \"Comment\") VALUES($1, $2, $3, $4, $5) RETURNING id"

const SELECT_EXPENSE_BY_DATE = "SELECT * FROM tx_expenses WHERE \"Date\" = $1"
