package main

import (
	"elelequent/prototypes/budget-api/dao/factory"
	"elelequent/prototypes/budget-api/dao/interfaces"
	"elelequent/prototypes/budget-api/utility"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var dao interfaces.BudgetDao

func main() {
	utility.LoadEnv()

	dao = factory.FactoryDao(os.Getenv("DB_ENGINE"))
	dao.EstablishConnection()

	http.HandleFunc("/expense/date/", expenseByDateHandler)

	fmt.Println("Starting Budget-API, ctrl+c to exit...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func expenseByDateHandler(w http.ResponseWriter, r *http.Request) {
	expenseDate := r.URL.Path[len("/expense/date/"):]

	expenses, err := dao.ExpensesByDate(expenseDate)

	if err != nil {
		log.Fatalf("Error getting expense for date %s: %s", expenseDate, err)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	expensesJson, err := json.Marshal(expenses)
	if err != nil {
		log.Fatalf("JSON error: %s", err)
	}

	fmt.Fprintf(w, "{\"Expenses\": %v}", string(expensesJson))
}
