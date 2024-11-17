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

func main() {
    log.Println("Starting application on AWS Lambda...")
    lambda.Start(Handler)
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return events.APIGatewayProxyResponse{
        StatusCode: 200,
        Headers: map[string]string{
            "Content-Type":                 "application/json",
            "Access-Control-Allow-Origin":  "*",
            "Access-Control-Allow-Methods": "GET, POST, PUT, DELETE, OPTIONS",
        },
        MultiValueHeaders: map[string][]string{
            "Set-Cookie": {"session=abc123; Path=/; HttpOnly", "user=john_doe; Path=/"},
        },
        Body: `{"message": "Hello, World!", "data": {"id": 1, "name": "Example"}}`,
        IsBase64Encoded: false,
    }, nil
}

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



