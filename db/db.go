package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConectaComBancoDeDados() *sql.DB {

	err := godotenv.Load("Properties/db_connection.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	conexao := fmt.Sprintf("user=%s dbname=%s password=%s host=localhost sslmode=disable", dbUser, dbName, dbPass)

	db, err := sql.Open("postgres", conexao)

	if err != nil {
		panic(err.Error())
	}

	return db
}
