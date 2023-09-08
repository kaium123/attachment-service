package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
	"github.com/spf13/viper"
)

var db *sql.DB

func InitDB() *sql.DB {
	// Open a PostgreSQL database connection (replace with your own connection string)
	dbUrl := viper.GetString("DB_URL")
	var err error
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Create the user table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS attachments (
			id SERIAL PRIMARY KEY,
			path TEXT,
			name TEXT,
			sourceType TEXT,
			sourceId INTEGER
		)
	`)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
