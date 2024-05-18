package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func (db *Database) Connect() error {
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

func (db *Database) ConnectionStatus() bool {
	err := db.connection.Ping()
	if err != nil {
		return false
	}
	return true
}
