package controllers

import (
	"const/core/orm/models"
	services "const/core/services/transaction"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/v4/types"
)

type TransactionController struct {
	transactionService *services.TransactionService
}

func NewTransactionController(transactionService *services.TransactionService) *TransactionController {
	return &TransactionController{transactionService: transactionService}
}

type Error struct {
	Message string `json:"message"`
}

type Transaction = models.Transacao

// @summary Purchase a ticket
// @description Purchase a ticket
// @tags transactions
// @accept json
// @produce json
// @param transaction body Transaction true "Transaction object"
// @success 201 {object} Transaction
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /transaction [post]
func (c *TransactionController) PurchaseTicket(ctx *gin.Context) {
	var purchaseRequest struct {
		TicketID     int           `json:"ticket_id"`
		CompradorID  int           `json:"comprador_id"`
		PrecoDeVenda types.Decimal `json:"preco_de_venda"`
	}

	if err := ctx.ShouldBindJSON(&purchaseRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction := &models.Transacao{
		Iddoticket:    purchaseRequest.TicketID,
		Iddocomprador: purchaseRequest.CompradorID,
		Precodevenda:  purchaseRequest.PrecoDeVenda,
	}

	if err := c.transactionService.PurchaseTicket(ctx, transaction); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, transaction)
}

// @summary Request refund
// @description Request refund
// @tags transactions
// @accept json
// @produce json
// @param id path int true "Transaction ID"
// @success 200
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /transaction/{id}/refund [put]
func (c *TransactionController) RequestRefund(ctx *gin.Context) {
	transactionIDStr := ctx.Param("id")
	transactionID, err := strconv.Atoi(transactionIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}

	if err := c.transactionService.RequestRefund(ctx, transactionID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func Handler(router *gin.RouterGroup, transactionService *services.TransactionService) {
	controller := NewTransactionController(transactionService)

	router.POST("/transaction", controller.PurchaseTicket)
	router.PUT("/transaction/:id/refund", controller.RequestRefund)
}
