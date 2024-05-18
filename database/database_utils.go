package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func (db *Database) connect() error {
	var err error
	db.connection, err = sql.Open("sqlite3", db.dsn)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (db *Database) Close() error {
	err := db.connection.Close()
	if err != nil {
		return err
	}

	return nil
}
