package database

import (
	"database/sql"
	"fmt"
	_ "log"

	_ "github.com/mattn/go-sqlite3"
)

func (db *Database) Connect() error {
	var err error
	db.connection, err = sql.Open("sqlite3", db.dsn)
	if err != nil {
		return err
	}

	ok, err := db.ConnectionStatus()
	if err != nil || !ok {
		return fmt.Errorf("error connecting to the database: %s", err)
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

func (db *Database) ConnectionStatus() (ok bool, err error) {
	err = db.connection.Ping()
	if err != nil {
		return false, err
	}

	return true, nil
}
