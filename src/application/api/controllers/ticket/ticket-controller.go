package controllers

import (
	"const/core/orm/models"
	services "const/core/services/ticket"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Ticket = models.Ticket
type TicketController struct {
	ticketService *services.TicketService
}
type Error struct {
	Message string `json:"message"`
}

func NewTicketController(ticketService *services.TicketService) *TicketController {
	return &TicketController{ticketService: ticketService}
}

// CreateTicket godoc
// @Summary Create a new ticket
// @Description Create a new ticket
// @Tags tickets
// @Accept json
// @Produce json
// @Param ticket body Ticket true "Ticket object"
// @Success 201 {object} Ticket
// @Failure 400 {object} Error
// @Failure 500 {object} Error
// @Router /tickets [post]
func (c *TicketController) CreateTicket(ctx *gin.Context) {
	var ticket models.Ticket
	if err := ctx.ShouldBindJSON(&ticket); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

// @summary Get available tickets by event
// @description Get available tickets by event
// @tags tickets
// @accept json
// @produce json
// @param eventID path int true "Event ID"
// @success 200 {array} Ticket
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /events/{eventID}/tickets [get]
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

// @summary Get tickets by seller
// @description Get tickets by seller
// @tags tickets
// @accept json
// @produce json
// @param userID path int true "User ID"
// @success 200 {array} Ticket
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /users/tickets/{userID} [get]
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

// @summary Mark ticket as used
// @description Mark ticket as used
// @tags tickets
// @accept json
// @produce json
// @param ticketID path int true "Ticket ID"
// @success 200
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /tickets/{ticketID}/use [put]
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

// @summary Authenticate ticket
// @description Authenticate ticket
// @tags tickets
// @accept json
// @produce json
// @param ticketID path int true "Ticket ID"
// @success 200
// @failure 400 {object} Error
// @failure 404 {object} Error
// @failure 409 {object} Error
// @failure 500 {object} Error
// @router /tickets/authenticate [post]
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
	return "UNIQUE-CODE-1234"
}

func Handler(router *gin.RouterGroup, ticketService *services.TicketService) {
	controller := NewTicketController(ticketService)

	router.POST("/tickets", controller.CreateTicket)
	router.GET("/events/:eventID/tickets", controller.GetAvailableTicketsByEvent)
	router.GET("/users/tickets/:userID", controller.GetTicketsBySeller)
	router.PUT("/tickets/:ticketID/use", controller.MarkTicketAsUsed)
	router.POST("/tickets/authenticate", controller.AuthenticateTicket)
}
