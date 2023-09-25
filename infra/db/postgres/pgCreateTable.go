package PGConfig

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func CreateDatabase(dbName string) (string, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("failed to connect .env")
		return "", err
	}
	db_dsn := os.Getenv("PG_DB_DSN")
	db_dsn_con := fmt.Sprintf("%s?sslmode=disable", db_dsn)
	db, err := sql.Open("postgres", db_dsn_con)

	if err != nil {
		log.Fatal(err)
	}
	createDatabaseQuery := fmt.Sprintf("SELECT datname FROM pg_database WHERE datname = '%s'", dbName)

	var existingDatabaseName string
	if err = db.QueryRow(createDatabaseQuery).Scan(&existingDatabaseName); err != nil {
		if err == sql.ErrNoRows {
			createDatabaseQuery := fmt.Sprintf("CREATE DATABASE %s", dbName)
			_, err := db.Exec(createDatabaseQuery)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
	new_db_dsn := fmt.Sprintf("%s%s?sslmode=disable", db_dsn, dbName)
	log.Println("Database Created!")
	return new_db_dsn, nil
}
func CreateTable(db *sql.DB) error {
	sqlFile := "infra/db/postgres/data.sql"
	sqlBytes, err := os.ReadFile(sqlFile)
	if err != nil {
		return err
	}
	sqlScript := string(sqlBytes)

	fmt.Println(sqlScript)

	_, err = db.Exec(sqlScript)
	if err != nil {
		if err.Error()[len(err.Error())-14:len(err.Error())] == "already exists" {
			log.Println("table already exists")
			return nil
		}
		return err
	}
	log.Println("Table Created!")
	return nil
}
