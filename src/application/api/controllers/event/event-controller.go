package controllers

import (
	"const/core/orm/models"
	services "const/core/services/event"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EventController struct {
	eventService *services.EventoService
}

func NewEventController(eventService *services.EventoService) *EventController {
	return &EventController{eventService: eventService}
}

func (c *EventController) CreateEvent(ctx *gin.Context) {
	var event models.Evento
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.eventService.CreateEvento(ctx, &event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

func (c *EventController) GetEventByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do evento inválido"})
		return
	}

	event, err := c.eventService.GetEventoByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Evento não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (c *EventController) GetAllEvents(ctx *gin.Context) {
	events, err := c.eventService.GetAllEventos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func (c *EventController) UpdateEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do evento inválido"})
		return
	}

	var event models.Evento
	if err := ctx.ShouldBindJSON(&event); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	event.Eventoid = id

	if err := c.eventService.UpdateEvento(ctx, &event); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, event)
}

func (c *EventController) DeleteEvent(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do evento inválido"})
		return
	}

	if err := c.eventService.DeleteEvento(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *EventController) GetEventsByTenant(ctx *gin.Context) {
	tenantIDStr := ctx.Param("tenantID")
	tenantID, err := strconv.Atoi(tenantIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do tenant inválido"})
		return
	}

	events, err := c.eventService.GetEventosByTenant(ctx, tenantID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, events)
}

func Handler(router *gin.RouterGroup, eventService *services.EventoService) {
	controller := NewEventController(eventService)

	router.POST("/event", controller.CreateEvent)
	router.GET("/event/:id", controller.GetEventByID)
	router.GET("/event", controller.GetAllEvents)
	router.PUT("/event/:id", controller.UpdateEvent)
	router.DELETE("/event/:id", controller.DeleteEvent)
	router.GET("/tenant/:tenantID/event", controller.GetEventsByTenant)
}
