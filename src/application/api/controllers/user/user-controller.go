package controllers

import (
	"const/core/orm/models"
	finanseServices "const/core/services/finance"
	userServices "const/core/services/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService    *userServices.UserService
	financeService *finanseServices.FinanceService
}

func NewUserController(userService *userServices.UserService, financeService *finanseServices.FinanceService) *UserController {
	return &UserController{
		userService:    userService,
		financeService: financeService}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user models.Usuario
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, _ := c.userService.GetUserByEmail(ctx, user.Email)
	if existingUser != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "E-mail já está em uso"})
		return
	}

	if err := c.userService.CreateUser(ctx, &user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetUserByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
		return
	}

	user, err := c.userService.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
		return
	}

	var userUpdates models.Usuario
	if err := ctx.ShouldBindJSON(&userUpdates); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUser, err := c.userService.GetUserByID(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	existingUser.Nome = userUpdates.Nome
	existingUser.Email = userUpdates.Email
	existingUser.Tenantid = userUpdates.Tenantid

	if err := c.userService.UpdateUser(ctx, id, existingUser); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, existingUser)
}

func (c *UserController) DeleteUser(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
		return
	}

	if err := c.userService.DeleteUser(ctx, id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (c *UserController) ListUsers(ctx *gin.Context) {
	users, err := c.userService.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetUserBalance(ctx *gin.Context) {
	userIDStr := ctx.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
		return
	}

	balance, err := c.financeService.GetUserBalance(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"balance": balance})
}

func (c *UserController) CreateUserNotificationPreferences(ctx *gin.Context) {
	var userNotificationPreferences models.Preferenciasdenotificacao
	if err := ctx.ShouldBindJSON(&userNotificationPreferences); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := c.userService.CreateNotificationPreferences(ctx, &userNotificationPreferences); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	user, err := c.userService.GetUserByID(ctx, userNotificationPreferences.Userid)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userNotificationPreferences.Userid = user.Userid

	ctx.JSON(http.StatusCreated, userNotificationPreferences)
}

func Handler(router *gin.RouterGroup, userService *userServices.UserService, financeService *finanseServices.FinanceService) {
	controller := NewUserController(userService, financeService)

	router.POST("/users", controller.CreateUser)
	router.GET("/users/:id", controller.GetUserByID)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
	router.GET("/users", controller.ListUsers)
}
