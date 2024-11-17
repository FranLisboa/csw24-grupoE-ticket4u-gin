package main

import (
	"const/infrastructure/database"
    setup "const/application/api/setup"
	"log"
	
	"context"
	

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

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    return ginLambda.ProxyWithContext(ctx,req)
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



