package handlers

import (
	"encoding/json"
	"fmt"
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
