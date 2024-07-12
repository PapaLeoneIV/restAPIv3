package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func NewDB(dbDriver string, dbSource string) *sql.DB {
	fmt.Printf("Opening the database\n")
	db, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Printf("failed to open the database connection: %v\n", err)
		return nil
	}
	fmt.Printf("Pinging Database\n")

	createTable(db)

	err = db.Ping()
	if err != nil {
		db.Close()
		log.Printf("failed to ping the database: %v\n", err)
		return nil
	}
	fmt.Printf("Database Connected\n")

	return db
}

func createTable(db *sql.DB) {
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		subject VARCHAR(100) NOT NULL,
		body VARCHAR(4096) NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	);
`)
	if err != nil {
		log.Fatalf("Error creating Students table: %v", err)
	}
}
