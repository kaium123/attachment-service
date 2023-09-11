package db

import (
	"attachment/common/logger"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var db *sql.DB

func InitDB() *sql.DB {
	dbUrl := viper.GetString("DB_URL")
	fmt.Println(dbUrl, "  dfgsdf")
	var err error
	db, err = sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err, " fffffff")
	}

	err = db.Ping()
	if err != nil {
		logger.LogError(err)
		log.Fatal(err)
	}

	// Create the "attachments" table
	err = createAttachmentsTable()
	if err != nil {
		logger.LogError(err)
		log.Fatal(err)
	}

	return db
}

func createAttachmentsTable() error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS attachments (
			id SERIAL PRIMARY KEY,
			path TEXT,
			name TEXT,
			sourceType TEXT,
			sourceId INTEGER
		)
	`)
	if err != nil {
		logger.LogError(err)
		return fmt.Errorf("error creating attachments table: %v", err)
	}
	return nil
}
