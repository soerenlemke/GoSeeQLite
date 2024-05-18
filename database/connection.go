package database

import (
	"database/sql"
	_ "log"

	_ "github.com/mattn/go-sqlite3"
)

func (db *Database) Connect() error {
	var err error
	db.connection, err = sql.Open("sqlite3", db.dsn)
	if err != nil {
		return err
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

func (db *Database) ConnectionStatus() (err error, ok bool) {
	err = db.connection.Ping()
	if err != nil {
		return err, false
	}

	return nil, true
}
