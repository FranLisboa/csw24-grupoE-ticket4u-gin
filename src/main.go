package main

import (
	setup "const/application/api/setup"
	"const/infrastructure/database"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "const/docs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda
var db *sql.DB

func david() {
	router := gin.Default()

	db := database.StartDB()
	defer db.Close()

	router.GET("/api/v1/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	setup.Setup(router, db)

	ginLambda = ginadapter.New(router)
}

func checkDBConnection() error {
	if db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	return db.Ping()
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if err := checkDBConnection(); err != nil {
		david()
		log.Printf("Database connection error: %v. Attempting to reconnect...", err)
		_, err := sql.Open("postgres", "postgresql://postgres:xyV4YBeY8Qz2FuZ@postgres.cy3myhw5bsdp.us-east-1.rds.amazonaws.com:5432/postgres")
		if err != nil {
			return events.APIGatewayProxyResponse{
				StatusCode: 500,
				Body:       "Internal Server Error: Database connection failed",
			}, nil
		}
	}
	return ginLambda.Proxy(req)
}

func main() {
	log.Println("Starting application on AWS Lambda...")

	lambda.Start(Handler)
}
