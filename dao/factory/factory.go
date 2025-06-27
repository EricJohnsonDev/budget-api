package factory

import (
	"elelequent/prototypes/budget-api/dao/interfaces"
	"elelequent/prototypes/budget-api/dao/postgresql"
	"log"
)

func FactoryDao(dbEngine string) interfaces.BudgetDao {
	var iface interfaces.BudgetDao

	switch dbEngine {
	case "postgresql":
		iface = postgresql.BudgetImplPostgresql{}
	default:
		log.Fatalf("DB %s is not implemented", dbEngine)
		return nil
	}

	return iface
}
