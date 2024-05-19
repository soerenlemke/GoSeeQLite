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

type TableColumn struct {
	Id               int
	Name             string
	Type             string
	NotNull          int
	DefaultValue     sql.NullString
	PrimaryKey       int
	ForeignKey       bool
	ReferencedTable  string
	ReferencedColumn string
}
