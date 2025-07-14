package handlers

import (
	"elelequent/prototypes/budget-api/dao/interfaces"
)

var dao interfaces.BudgetDao

func SetDao(daoIn interfaces.BudgetDao) {
	dao = daoIn
}
