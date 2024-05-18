package database

import "fmt"

func NewDatabase(pathToDb string) (Database, error) {
	if pathToDb == "" {
		return Database{}, fmt.Errorf("path to db can't be empty")
	}

	d := Database{
		dsn: pathToDb,
	}
	d.Get = &Get{DB: &d}
	d.Set = &Set{DB: &d}

	err := d.Connect()
	if err != nil {
		return Database{}, fmt.Errorf("can`t connect to database: %s", d.dsn)
	}

	return d, err
}
