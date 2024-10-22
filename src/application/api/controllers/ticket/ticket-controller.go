package controllers

import (
	"const/core/orm/models"
	services "const/core/services/ticket"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TicketController struct {
	ticketService *services.TicketService
}

func NewTicketController(ticketService *services.TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

func (c *TicketController) CreateTicket(ctx *gin.Context) {
	var ticket models.Ticket
	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// assume que o id do vendedor vem do body da requisição
	if ticket.Iddovendedor == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Seller ID is required"})
		return
	}

	ticket.Codigounicodeverificacao = generateUniqueCode()

	ticket.Status = "disponivel"

	if err := c.ticketService.CreateTicket(ctx, &ticket); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, ticket)
}

func (c *TicketController) GetAvailableTicketsByEvent(ctx *gin.Context) {
	eventIDStr := ctx.Param("eventID")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	tickets, err := c.ticketService.GetAvailableTicketsByEvent(ctx, eventID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (c *TicketController) GetTicketsBySeller(ctx *gin.Context) {
	sellerIDStr := ctx.Param("userID")
	sellerID, err := strconv.Atoi(sellerIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid seller ID"})
		return
	}

	tickets, err := c.ticketService.GetTicketsBySeller(ctx, sellerID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tickets)
}

func (c *TicketController) MarkTicketAsUsed(ctx *gin.Context) {
	ticketIDStr := ctx.Param("ticketID")
	ticketID, err := strconv.Atoi(ticketIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}

	if err := c.ticketService.MarkTicketAsUsed(ctx, ticketID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *TicketController) AuthenticateTicket(ctx *gin.Context) {
	var request struct {
		CodigoUnicoDeVerificacao string `json:"codigo_unico_de_verificacao"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.ticketService.AuthenticateTicket(ctx, request.CodigoUnicoDeVerificacao)
	if err != nil {
		switch err.Error() {
		case "ingresso não encontrado":
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case "ingresso já foi utilizado":
			ctx.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Ingresso autenticado com sucesso"})
}

func generateUniqueCode() string {
	return uuid.New().String()
}

func Handler(router *gin.RouterGroup, ticketService *services.TicketService) {
	controller := NewTicketController(ticketService)

	router.POST("/tickets", controller.CreateTicket)
	router.GET("/events/:eventID/tickets", controller.GetAvailableTicketsByEvent)
	router.GET("/users/tickets/:userID", controller.GetTicketsBySeller)
	router.PUT("/tickets/:ticketID/use", controller.MarkTicketAsUsed)
	router.POST("/tickets/authenticate", controller.AuthenticateTicket)
}
