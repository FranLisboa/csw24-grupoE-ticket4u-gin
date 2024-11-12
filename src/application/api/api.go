package api

import (
	eventController "const/application/api/controllers/event"
	feedbackController "const/application/api/controllers/feedback"
	tenantController "const/application/api/controllers/tenant"
	ticketController "const/application/api/controllers/ticket"
	transactionController "const/application/api/controllers/transaction"
	userController "const/application/api/controllers/user"

	event "const/core/services/event"
	feedback "const/core/services/feedback"
	finance "const/core/services/finance"
	tenant "const/core/services/tenant"
	ticket "const/core/services/ticket"
	transaction "const/core/services/transaction"
	user "const/core/services/user"

	"database/sql"
	"net/http"

	_ "const/docs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var ginLambda *ginadapter.GinLambda

func Setup(router *gin.Engine, db *sql.DB) {

	userService := user.NewService(db)
	eventService := event.NewEventoService(db)
	tenantService := tenant.NewTenantService(db)
	financeService := finance.NewFinanceService(db)
	ticketService := ticket.NewTicketService(db)
	transactionService := transaction.NewTransactionService(db, ticketService, financeService)
	feedbackService := feedback.NewFeedbackService(db)

	v1 := router.Group("/api/v1")

	userController.Handler(v1, &userService, financeService)
	tenantController.Handler(v1, tenantService)
	eventController.Handler(v1, eventService)
	ticketController.Handler(v1, ticketService)
	transactionController.Handler(v1, transactionService)
	feedbackController.Handler(v1, feedbackService)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func Init(db *sql.DB) {
	router := gin.Default()

	Setup(router, db)

	ginLambda = ginadapter.New(router)
	main()
}

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return ginLambda.Proxy(req)
}

func main() {

	lambda.Start(Handler)
}
