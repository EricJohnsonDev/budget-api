package handlers

import (
	"elelequent/prototypes/budget-api/dao/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func ExpensesByDates(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	startDate := r.Form.Get("start")
	endDate := r.Form.Get("end")

	if !isValidDateFormat(startDate) {
		http.Error(w, "Start date is required, but is either missing or not a valid format (expected MM-dd-YYYY)", http.StatusBadRequest)
		return
	}

	if !isValidDateFormat(endDate) {
		http.Error(w, "End date is missing or not a valid format (expected MM-dd-YYYY).  Restricting request to the start date.", http.StatusContinue)
		endDate = startDate
	}

	expenses, err := dao.ExpensesByDate(startDate, endDate)

	if err != nil {
		log.Printf("ERROR Unable to retrieve expenses with start %s and end %s: %v", startDate, endDate, err)
		http.Error(w, "An error occurred while retrieving data", http.StatusInternalServerError)
		return
	}

	expensesJson, err := json.Marshal(expenses)
	if err != nil {
		log.Printf("ERROR Unable to encode JSON: %v", err)
		http.Error(w, "An error occurred while processing data", http.StatusInternalServerError)
	}

	setCommonHeaders(w)

	fmt.Fprintf(w, "{\"Expenses\": %v}", string(expensesJson))
}

func AddExpenses(w http.ResponseWriter, r *http.Request) {
	var expensesToAdd []models.Tx_expenses

	setCommonHeaders(w)

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	err = json.Unmarshal(body, &expensesToAdd)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Error decoding json body", http.StatusInternalServerError)
		return
	}

	numRowsAdded, err := dao.AddExpenses(expensesToAdd)
	if err != nil {
		log.Printf("Error: %v", err)
		http.Error(w, "Error adding expenses", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, numRowsAdded)
}
