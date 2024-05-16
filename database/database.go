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

func NewDatabase(pathToDb string) (Database, error) {
	if pathToDb == "" {
		return Database{}, fmt.Errorf("path to db can't be empty")
	}

	d := Database{
		dsn: pathToDb,
	}

	err := d.connect()
	if err != nil {
		return Database{}, fmt.Errorf("can`t connect to database: %s", d.dsn)
	}

	return d, err
}

func (db *Database) connect() error {
	var err error
	db.connection, err = sql.Open("sqlite3", db.dsn)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (db *Database) GetAllTableNames() ([]string, error) {
	queryString := fmt.Sprintf("SELECT name FROM %s WHERE type='table';", db.dsn)
	rows, err := db.connection.Query(queryString)
	if err != nil {
		return nil, err
	}

	var tableNames []string
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tableNames = append(tableNames, tableName)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tableNames, nil
}
