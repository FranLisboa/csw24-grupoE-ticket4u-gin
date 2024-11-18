package main

import (
	"const/infrastructure/database"
    setup "const/application/api/setup"
	"log"	

	_ "const/docs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

var ginLambda *ginadapter.GinLambda

func init() {
	router := gin.Default()

	db := database.StartDB()
	defer db.Close()

	router.GET("/api/v1/health", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": "ok",
        })
    })

	setup.Setup(router, db)

	ginLambda = ginadapter.New(router)
}


func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return ginLambda.Proxy(req)
}

func main() {
    log.Println("Starting application on AWS Lambda...")

    lambda.Start(Handler)
}

