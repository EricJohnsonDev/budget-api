package postgresql

import (
	"database/sql"
	"elelequent/prototypes/budget-api/dao/models"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresqlDao struct {
}

var db *sql.DB

func (dao PostgresqlDao) EstablishConnection() {
	db = connectToDb()
}

func (dao PostgresqlDao) AddExpenses(expenses []models.Tx_expenses) (int64, error) {
	values := []interface{}{}
	numTxExpenseFields := 6

	for _, expense := range expenses {
		values = append(values, expense.Date, expense.Amount, expense.Institution, expense.Category, expense.Subcategory, expense.Comment)
	}

	stmt, err := db.Prepare(INSERT_EXPENSES + prepareValuesFmt(numTxExpenseFields, len(expenses)))
	if err != nil {
		log.Printf("Error: Unable to prepare insert: %v", err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		log.Printf("Error: Unable to execute prepared insert: %v", err)
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()

	return rowsAffected, err
}

func (dao PostgresqlDao) ExpensesByDate(startDate string, endDate string) ([]models.Tx_expenses, error) {
	var expenses []models.Tx_expenses

	rows, err := db.Query(SELECT_EXPENSE_BY_DATE, startDate, endDate)
	if err != nil {
		return nil, fmt.Errorf("expensesByDate %q - %q: %v", startDate, endDate, err)
	}
	defer rows.Close()

	for rows.Next() {
		var expense models.Tx_expenses
		if err := rows.Scan(&expense.ID, &expense.Date, &expense.Amount, &expense.Institution, &expense.Category, &expense.Subcategory, &expense.Comment); err != nil {
			return nil, fmt.Errorf("expensesByDate %q - %q: %v", startDate, endDate, err)
		}

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("expensesByDate %q - %q: %v", startDate, endDate, err)
	}

	return expenses, nil
}
