package postgresql

import (
	"database/sql"
	"elelequent/prototypes/budget-api/dao/models"
	"fmt"

	_ "github.com/lib/pq"
)

type PostgresqlDao struct {
}

var db *sql.DB

func (dao PostgresqlDao) EstablishConnection() {
	db = connectToDb()
}

func (dao PostgresqlDao) AddExpense(expense models.Tx_expenses) (int, error) {
	row := db.QueryRow(INSERT_EXPENSE, expense.Date, expense.Amount, expense.Institution, expense.Category, expense.Comment)
	err := row.Scan(&expense.ID)

	if err != nil {
		return 0, fmt.Errorf("addExpense: %v", err)
	}

	return expense.ID, nil
}

func (dao PostgresqlDao) ExpensesByDate(date string) ([]models.Tx_expenses, error) {
	var expenses []models.Tx_expenses

	rows, err := db.Query(SELECT_EXPENSE_BY_DATE, date)
	if err != nil {
		return nil, fmt.Errorf("expensesByDate %q: %v", date, err)
	}
	defer rows.Close()

	for rows.Next() {
		var expense models.Tx_expenses
		if err := rows.Scan(&expense.ID, &expense.Date, &expense.Amount, &expense.Institution, &expense.Category, &expense.Subcategory, &expense.Comment); err != nil {
			return nil, fmt.Errorf("expensesByDate %q: %v", date, err)
		}

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("expensesByDate %q: %v", date, err)
	}

	return expenses, nil
}
