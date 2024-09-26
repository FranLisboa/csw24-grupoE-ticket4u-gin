package main

import (
	middleware "./api/middlewares"
	"./infrastructure/database"
	"./src/api"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.StartDB()
	defer db.Close()

	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	myApp := api.NewApp(db)

	port := os.Getenv("API_PORT")
	if port == "" {
		port = "8080"
	}

	api.RunServer(port, myApp)
}
