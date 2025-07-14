package handlers

import (
	"elelequent/prototypes/budget-api/dao/interfaces"
	"net/http"
	"regexp"
)

var dao interfaces.BudgetDao

func SetDao(daoIn interfaces.BudgetDao) {
	dao = daoIn
}

func setCommonHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func isValidDateFormat(dateIn string) bool {
	return regexp.MustCompile(`\d{2}-\d{2}-\d{4}`).MatchString(dateIn)
}
