package repository

import (
	"database/sql"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {

	envFile := ".env"

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDbName := os.Getenv("DB_NAME")

	// Connect to MySQL
	db, err := sql.Open("mysql", dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDbName + "?parseTime=true")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return db, nil
}