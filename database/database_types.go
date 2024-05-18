package database

import "database/sql"

type Database struct {
	dsn        string
	connection *sql.DB
	Get        *Get
	Set        *Set
}

type Get struct {
	DB *Database
}

type Set struct {
	DB *Database
}
