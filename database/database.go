package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	dsn        string
	connection *sql.DB
}

func (db *Database) NewDatabase(pathToDb string) error {
	if pathToDb != "" {
		db.dsn = pathToDb
	} else {
		return fmt.Errorf("path to db can't be empty")
	}

	err := db.testConnection()
	if err != nil {
		return fmt.Errorf("can't connect to database: %s", db.dsn)
	}

	return nil
}

func (db *Database) testConnection() error {
	var err error
	db.connection, err = sql.Open("sqlite3", db.dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.connection.Close()

	return nil
}
