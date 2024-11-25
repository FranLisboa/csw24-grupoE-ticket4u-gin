package main

import (
	"const/application/api"
	"const/infrastructure/database"
	"log"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.StartDB()
	defer db.Close()

	if os.Getenv("RUN_MODE") == "local" {

		gin.Default()

		myApp := api.NewApp(db)

		port := os.Getenv("API_PORT")
		if port == "" {
			port = "8080"
		}

		api.RunServer(port, myApp)

	} else {
		log.Println("Starting application on AWS Lambda...")
		api.Init(db)
	}
}
