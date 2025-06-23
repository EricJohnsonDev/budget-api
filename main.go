package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var db *sql.DB

type ExpenseTx struct {
	ID          int
	Date        string
	Amount      string
	Institution string
	Category    string
	Subcategory string
	Comment     string
}

func main() {
	loadEnv()
	db = connectToDb()

	expenses, err := expensesByDate("2025-05-29")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Expenses found: %v\n", expenses)
}

// Query for expenses from a single day
func expensesByDate(date string) ([]ExpenseTx, error) {
	var expenses []ExpenseTx

	rows, err := db.Query("SELECT * FROM tx_expenses WHERE \"Date\" = $1", date)
	if err != nil {
		return nil, fmt.Errorf("expensesByDate %q: %v", date, err)
	}
	defer rows.Close()

	for rows.Next() {
		var expense ExpenseTx
		if err := rows.Scan(&expense.ID, &expense.Date, &expense.Amount, &expense.Institution, &expense.Category, &expense.Subcategory, &expense.Comment); err != nil {
			return nil, fmt.Errorf("expensesByDate %q: %v", date, err)
		}

		expenses = append(expenses, expense)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("expensesByDate %q: %v", date, err)
	}

	return expenses, nil
}

func connectToDb() *sql.DB {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbPwd := os.Getenv("DB_PWD")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	sslCert := os.Getenv("SSL_CERT")

	// Database stuff
	connStr := "host=" + dbHost +
		" port=" + dbPort +
		" user=" + dbUser +
		" password=" + dbPwd +
		" dbname=" + dbName +
		" sslmode=verify-full " +
		" sslrootcert=" + sslCert

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Confirmed connection to db")
	return db
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
