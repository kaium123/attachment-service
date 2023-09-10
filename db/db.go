package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" 
	"github.com/spf13/viper"
)

var db *sql.DB

func InitDB() *sql.DB {
	dbUrl := viper.GetString("DB_URL")
	var err error
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

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
