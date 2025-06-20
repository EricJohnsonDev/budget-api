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

func main() {
	loadEnv()

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

	fmt.Println("Connected!")
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
