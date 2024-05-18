package database

import (
	"database/sql"
	"fmt"
	"log"
	"path/filepath"
)

func (g *Get) DatabaseName() (string, error) {
	databaseName := filepath.Base(g.DB.dsn)
	if databaseName == "" {
		return "", fmt.Errorf("error getting the name of the given database")
	}

	return databaseName, nil
}

func (g *Get) AllTableNames() ([]string, error) {
	queryString := "SELECT name FROM sqlite_master WHERE type='table';"
	rows, err := g.DB.connection.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows:", err)
		}
	}(rows)

	var tableNames []string
	for rows.Next() {
		var tableName string
		err := rows.Scan(&tableName)
		if err != nil {
			return nil, err
		}
		tableNames = append(tableNames, tableName)
	}

	return tableNames, nil
}

func (g *Get) TableColumns(t string) ([]ColumnInfo, error) {
	queryString := fmt.Sprintf("PRAGMA table_info(%s);", t)
	columns, err := g.DB.connection.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing columns:", err)
		}
	}(columns)

	var result []ColumnInfo
	for columns.Next() {
		var tc TableColumn
		err := columns.Scan(&tc.Id, &tc.Name, &tc.Type, &tc.NotNull, &tc.DefaultValue, &tc.PrimaryKey)
		if err != nil {
			return nil, err
		}
		result = append(result, ColumnInfo{Name: tc.Name, Type: tc.Type})
	}

	if err = columns.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
