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

func (g *Get) TableColumns(t string) ([]TableColumn, error) {
	rows, err := g.DB.connection.Query(fmt.Sprintf("PRAGMA table_info(%s);", t))
	if err != nil {
		return nil, err
	}

	defer func() {
		err := rows.Close()
		if err != nil {
			log.Println("Error closing rows: ", err)
		}
	}()

	var result []TableColumn
	for rows.Next() {
		var tc TableColumn
		err := rows.Scan(&tc.Id, &tc.Name, &tc.Type, &tc.NotNull, &tc.DefaultValue, &tc.PrimaryKey)
		if err != nil {
			return nil, err
		}
		result = append(result, TableColumn{Name: tc.Name, Type: tc.Type})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}
