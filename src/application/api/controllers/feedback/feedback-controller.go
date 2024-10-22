package controllers

import (
	"const/core/orm/models"

	services "const/core/services/feedback"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Avaliacao = models.Avaliacao
type FeedbackController struct {
	avaliacaoService *services.FeedbackService
}

type Error struct {
	Message string `json:"message"`
}

func NewFeedbackController(avaliacaoService *services.FeedbackService) *FeedbackController {
	return &FeedbackController{avaliacaoService: avaliacaoService}
}

// @summary Create a new feedback
// @description Create a new feedback
// @tags feedback
// @accept json
// @produce json
// @param avaliacao body Avaliacao true "Avaliacao object"
// @success 201 {object} Avaliacao
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /feedback [post]
func (c *FeedbackController) CreateAvaliacao(ctx *gin.Context) {
	var avaliacao models.Avaliacao
	if err := ctx.ShouldBindJSON(&avaliacao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.avaliacaoService.CreateFeedback(ctx, &avaliacao); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, avaliacao)
}

// @summary Get feedbacks by seller
// @description Get feedbacks by seller
// @tags feedback
// @accept json
// @produce json
// @param vendedorID path int true "Seller ID"
// @success 200 {array} Avaliacao
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /vendedor/{vendedorID}/feedback [get]
func (c *FeedbackController) GetAvaliacoesByVendedor(ctx *gin.Context) {
	vendedorIDStr := ctx.Param("vendedorID")
	vendedorID, err := strconv.Atoi(vendedorIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do vendedor inválido"})
		return
	}

	avaliacoes, err := c.avaliacaoService.GetFeedbacksByVendedor(ctx, vendedorID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, avaliacoes)
}

func Handler(router *gin.RouterGroup, avaliacaoService *services.FeedbackService) {
	controller := NewFeedbackController(avaliacaoService)

	router.POST("/avaliacao", controller.CreateAvaliacao)
	router.GET("/vendedor/:vendedorID/avaliacao", controller.GetAvaliacoesByVendedor)
}
