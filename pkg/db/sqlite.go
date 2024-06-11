package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) *sql.DB {
	db, err := sql.Open("sqlite3", filepath)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS metrics (
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"timestamp" DATETIME DEFAULT CURRENT_TIMESTAMP,
		"metric_name" TEXT,
		"value" REAL
	);`

	if _, err := db.Exec(createTableSQL); err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	return db
}
