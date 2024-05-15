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

	db := &database.Database{}
	err = db.NewDatabase(pathToDb)
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}

	fmt.Println("Successfully connected to the database")
}
