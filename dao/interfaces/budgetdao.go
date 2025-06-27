package interfaces

import "elelequent/prototypes/budget-api/dao/models"

type BudgetDao interface {
	AddExpense(expense models.Tx_expenses) (int, error)
	ExpensesByDate(date string) ([]models.Tx_expenses, error)
}
