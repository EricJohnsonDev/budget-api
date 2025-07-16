package postgresql

import (
	"strconv"
	"strings"
)

const INSERT_EXPENSES = "INSERT INTO tx_expenses(\"Date\", \"Amount\", \"Institution\", \"Category\", \"Subcategory\", \"Comment\") VALUES"

const SELECT_EXPENSE_BY_DATE = "SELECT * FROM tx_expenses WHERE \"Date\" BETWEEN $1 AND $2;"

func prepareValuesFmt(numValues, numRows int) string {
	var result string

	for i := 0; i < numRows; i++ {
		result += "("

		for j := 1; j <= numValues; j++ {
			result += "$" + strconv.Itoa(i*numValues+j) + ","
		}
		result = strings.Trim(result, ",")
		result += "),"
	}

	return strings.Trim(result, ",")
}
