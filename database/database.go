package database

import (
	"fmt"
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
