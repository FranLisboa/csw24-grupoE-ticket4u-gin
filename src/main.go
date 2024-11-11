package main

import (
	"const/application/api"
	"const/infrastructure/database"
	"log"
)

func main() {
	db := database.StartDB()
	defer db.Close()

	log.Println("Starting application on AWS Lambda...")
	api.Init(db)
}
