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
			fmt.Println("Error closing the database:", err)
			return
		}
	}(&db)

	fmt.Println("Successfully connected to the database")

	tableNames, err := db.GetAllTableNames()
	if err != nil {
		fmt.Println("Error getting the table names of the database:", err)
		return
	}
	fmt.Println(tableNames)
}
