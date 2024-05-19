package main

import (
	"GoSeeQLite/database"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pathToDb := os.Getenv("DATABASE_PATH")

	db, err := database.NewDatabase(pathToDb)
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer func(db *database.Database) {
		err := db.Close()
		if err != nil {
			fmt.Println("Error closing the database: ", err)
			return
		}
	}(&db)

	fmt.Println("Successfully connected to the database")

	databaseName, err := db.Get.DatabaseName()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(databaseName)

	tableNames, err := db.Get.AllTableNames()
	if err != nil {
		fmt.Println("Error getting the table names of the database: ", err)
		return
	}
	fmt.Println(tableNames)

	// Iterate over all tables and print column information
	for _, tableName := range tableNames {
		fmt.Println("Table:", tableName)
		columnInfos, err := db.Get.TableColumns(tableName)
		if err != nil {
			fmt.Println("Error getting the column names for the table", tableName, ":", err)
			return
		}
		for _, columnInfo := range columnInfos {
			fmt.Printf("Column ID: %d, Name: %s, Type: %s, Not Null: %d, Default Value: %v, Primary Key: %d\n",
				columnInfo.Id, columnInfo.Name, columnInfo.Type, columnInfo.NotNull, columnInfo.DefaultValue.String, columnInfo.PrimaryKey)
		}
		fmt.Println()
	}
}
