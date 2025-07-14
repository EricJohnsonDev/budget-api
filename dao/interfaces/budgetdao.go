package interfaces

import "elelequent/prototypes/budget-api/dao/models"

type BudgetDao interface {
	EstablishConnection()
	AddExpense(expense models.Tx_expenses) (int, error)
	ExpensesByDate(date string) ([]models.Tx_expenses, error)
}
