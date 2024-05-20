package database

import (
	"database/sql"
	_ "log"

	_ "github.com/mattn/go-sqlite3"
)

func (db *Database) Connect() error {
	// Does not really check if the database file 
	// is a valid database. This is done by ConnectionStatus()

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
	// Returns error if the file is not a valid database
	
	err = db.connection.Ping()
	if err != nil {
		return err, false
	}

	return nil, true
}
