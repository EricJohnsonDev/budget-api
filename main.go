package main

import (
	"elelequent/prototypes/budget-api/dao/factory"
	"elelequent/prototypes/budget-api/dao/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	loadEnv()

	budgetDao := factory.FactoryDao("postgresql")

	expenses, err := budgetDao.ExpensesByDate("2025-05-29")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Expenses found: %v\n", expenses)

	testExpense := models.Tx_expenses{
		Date:        "2000-01-01",
		Amount:      "$666.42",
		Institution: "Budget-API",
		Category:    "TESTING",
		Comment:     "Testing, sent from api",
	}
	newRowId, err := budgetDao.AddExpense(testExpense)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Added expense with ID: %v\n", newRowId)
}

func loadEnv() {
	env := os.Getenv("BUDGETAPI_ENV")
	if env == "" {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if env != "test" {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load()
}
