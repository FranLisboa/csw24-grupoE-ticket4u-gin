package api

import (
	eventController "const/application/api/controllers/event"
	tenantController "const/application/api/controllers/tenant"
	ticketController "const/application/api/controllers/ticket"
	transactionController "const/application/api/controllers/transaction"
	userController "const/application/api/controllers/user"

	event "const/core/services/event"
	finance "const/core/services/finance"
	tenant "const/core/services/tenant"
	ticket "const/core/services/ticket"
	transaction "const/core/services/transaction"
	user "const/core/services/user"

	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, db *sql.DB) {
	userService := user.NewService(db)
	eventService := event.NewEventoService(db)
	tenantService := tenant.NewTenantService(db)
	financeService := finance.NewFinanceService(db)
	ticketService := ticket.NewTicketService(db)
	transactionService := transaction.NewTransactionService(db, ticketService, financeService)

	v1 := router.Group("/api/v1")

	userController.Handler(v1, userService)
	tenantController.Handler(v1, tenantService)
	eventController.Handler(v1, eventService)
	ticketController.Handler(v1, ticketService)
	transactionController.Handler(v1, transactionService)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

}
