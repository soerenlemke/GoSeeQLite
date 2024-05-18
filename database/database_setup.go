package database

import (
	"database/sql"
	"fmt"
)

func NewDatabase(pathToDb string) (*Database, error) {
	if pathToDb == "" {
		return nil, fmt.Errorf("path to db can't be empty")
	}

	db, err := sql.Open("sqlite3", pathToDb)
	if err != nil {
		return nil, fmt.Errorf("can't connect to database: %s", pathToDb)
	}

	d := &Database{
		dsn:        pathToDb,
		connection: db,
	}

	d.Get = &Get{DB: d}

	return d, nil
}
