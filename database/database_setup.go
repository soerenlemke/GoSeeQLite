package database

import "fmt"

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
