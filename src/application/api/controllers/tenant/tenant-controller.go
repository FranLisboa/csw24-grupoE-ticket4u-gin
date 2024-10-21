package controllers

import (
	"const/core/orm/models"
	services "const/core/services/tenant"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TenantController struct {
	tenantService *services.TenantService
}

func NewTenantController(tenantService *services.TenantService) *TenantController {
	return &TenantController{tenantService: tenantService}
}

func (c *TenantController) CreateTenant(ctx *gin.Context) {
	var tenant models.Tenant
	if err := ctx.ShouldBindJSON(&tenant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.tenantService.CreateTenant(ctx, &tenant); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, tenant)
}

func (c *TenantController) GetTenantByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do tenant inválido"})
		return
	}

	tenant, err := c.tenantService.GetTenantByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Tenant não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, tenant)
}

func (c *TenantController) GetAllTenants(ctx *gin.Context) {
	tenants, err := c.tenantService.GetAllTenants(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tenants)
}

func (c *TenantController) UpdateTenant(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do tenant inválido"})
		return
	}

	var tenant models.Tenant
	if err := ctx.ShouldBindJSON(&tenant); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tenant.Tenantid = id

	if err := c.tenantService.UpdateTenant(ctx, &tenant); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tenant)
}

func (c *TenantController) DeleteTenant(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID do tenant inválido"})
		return
	}

	if err := c.tenantService.DeleteTenant(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func Handler(router *gin.RouterGroup, tenantService *services.TenantService) {
	controller := NewTenantController(tenantService)

	router.POST("/tenant", controller.CreateTenant)
	router.GET("/tenant/:id", controller.GetTenantByID)
	router.GET("/tenant", controller.GetAllTenants)
	router.PUT("/tenant/:id", controller.UpdateTenant)
	router.DELETE("/tenant/:id", controller.DeleteTenant)
}
