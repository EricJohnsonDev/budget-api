package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ExpensesByDates(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	startDate := r.Form.Get("start")
	endDate := r.Form.Get("end")

	w.Header().Set("Access-Control-Allow-Origin", "*")

	expenses, err := dao.ExpensesByDate(startDate, endDate)

	if err != nil {
		http.Error(w, "An error occurred while retrieving data", http.StatusInternalServerError)
		return
	}

	expensesJson, err := json.Marshal(expenses)
	if err != nil {
		http.Error(w, "An error occurred while processing data", http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "{\"Expenses\": %v}", string(expensesJson))
}
