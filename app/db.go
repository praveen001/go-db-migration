package app

import (
	"database/sql"
	"fmt"
)

// NewDB .
func NewDB() *sql.DB {
	fmt.Println("Connecting to MySQL database...")

	db, err := sql.Open("mysql", "root:welcome@tcp(127.0.0.1:3306)/migrationtest")
	if err != nil {
		fmt.Println("Unable to connect to database", err.Error())
		return nil
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Unable to connect to database", err.Error())
		return nil
	}

	fmt.Println("Database connected!")

	return db
}
