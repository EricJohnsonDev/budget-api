package postgresql

import (
	"elelequent/prototypes/budget-api/dao/models"
	"fmt"
)

type BudgetImplPostgresql struct {
}

func (dao BudgetImplPostgresql) AddExpense(expense models.Tx_expenses) (int, error) {
	db := connectToDb()
	defer db.Close()

	row := db.QueryRow("INSERT INTO tx_expenses (\"Date\", \"Amount\", \"Institution\", \"Category\", \"Comment\") VALUES($1, $2, $3, $4, $5) RETURNING id", expense.Date, expense.Amount, expense.Institution, expense.Category, expense.Comment)
	err := row.Scan(&expense.ID)

	if err != nil {
		return 0, fmt.Errorf("addExpense: %v", err)
	}

	return expense.ID, nil
}

func (dao BudgetImplPostgresql) ExpensesByDate(date string) ([]models.Tx_expenses, error) {
	var expenses []models.Tx_expenses

	db := connectToDb()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM tx_expenses WHERE \"Date\" = $1", date)
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
