package database

import (
	"fmt"
	"os"
)

func NewDatabase(pathToDb string) (Database, error) {
	if pathToDb == "" {
		return Database{}, fmt.Errorf("path to db can't be empty")
	}
	if _, err := os.Stat(pathToDb); os.IsNotExist(err){
		return Database{}, fmt.Errorf("file does not exist")
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

	err, ok := d.ConnectionStatus()
	if !ok {
		// The file is not a database
		return Database{}, err
	}

	return d, err
}
