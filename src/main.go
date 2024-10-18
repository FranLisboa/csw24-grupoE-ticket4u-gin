package main

import (
	"const/application/api"
	"const/infrastructure/database"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	db := database.StartDB()
	defer db.Close()

	gin.Default()

	myApp := api.NewApp(db)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	api.RunServer(port, myApp)
}
