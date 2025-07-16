package interfaces

import "elelequent/prototypes/budget-api/dao/models"

type BudgetDao interface {
	EstablishConnection()
	AddExpenses(expenses []models.Tx_expenses) (int64, error)
	ExpensesByDate(startDate string, endDate string) ([]models.Tx_expenses, error)
}
