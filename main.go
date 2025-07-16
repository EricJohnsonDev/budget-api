package main

import (
	"elelequent/prototypes/budget-api/dao/factory"
	"elelequent/prototypes/budget-api/handlers"
	"elelequent/prototypes/budget-api/utility"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	utility.LoadEnv()

	dao := factory.FactoryDao(os.Getenv("DB_ENGINE"))
	dao.EstablishConnection()
	handlers.SetDao(dao)

	http.HandleFunc("/expense/date", handlers.ExpensesByDates)
	http.HandleFunc("/expense/add", handlers.AddExpenses)

	fmt.Println("Starting Budget-API, ctrl+c to exit...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
