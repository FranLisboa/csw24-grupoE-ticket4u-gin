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

type Error struct {
	Message string `json:"message"`
}

type User = models.Usuario
type Balance struct {
	Balance float64 `json:"balance"`
}
type UserNotificationPreferences = models.Preferenciasdenotificacao

// @summary Create a new user
// @description Create a new user
// @tags users
// @accept json
// @produce json
// @param user body User true "User object"
// @success 201 {object} User
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /api/v1/users [post]
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

// @summary Get user by ID
// @description Get user by ID
// @tags users
// @accept json
// @produce json
// @param id path int true "User ID"
// @success 200 {object} User
// @failure 400 {object} Error
// @failure 404 {object} Error
// @router /api/v1/users/{id} [get]
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

// @summary Update user
// @description Update user
// @tags users
// @accept json
// @produce json
// @param id path int true "User ID"
// @param user body User true "User object"
// @success 200 {object} User
// @failure 400 {object} Error
// @failure 404 {object} Error
// @failure 500 {object} Error
// @router /api/v1/users/{id} [put]
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

// @summary Delete user
// @description Delete user
// @tags users
// @param id path int true "User ID"
// @success 204
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /api/v1/users/{id} [delete]
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

// @summary Get all users
// @description Get all users
// @tags users
// @accept json
// @produce json
// @success 200 {array} User
// @router /api/v1/users [get]
func (c *UserController) ListUsers(ctx *gin.Context) {
	users, err := c.userService.GetUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, users)
}

// @summary Get user balance
// @description Get user balance
// @tags users
// @accept json
// @produce json
// @param id path int true "User ID"
// @success 200 {object} Balance
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /api/v1/users/{id}/balance [get]
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

// @summary Create user notification preferences
// @description Create user notification preferences
// @tags users
// @accept json
// @produce json
// @param user body UserNotificationPreferences true "UserNotificationPreferences object"
// @success 201 {object} UserNotificationPreferences
// @failure 400 {object} Error
// @failure 500 {object} Error
// @router /api/v1/users/notification-preferences [post]
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

// @summary Update user notification preferences
// @description Update user notification preferences
// @tags users
// @accept json
// @produce json
// @param id path int true "User ID"
// @param user body UserNotificationPreferences true "UserNotificationPreferences object"
// @success 200 {object} UserNotificationPreferences
// @failure 400 {object} Error
// @failure 404 {object} Error
// @failure 500 {object} Error
// @router /api/v1/users/notification-preferences/{id} [put]
func (c *UserController) UpdateUserNotificationPreferences(ctx *gin.Context) {
	userIDStr := ctx.Param("id")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
		return
	}

	var userNotificationPreferences models.Preferenciasdenotificacao
	if err := ctx.ShouldBindJSON(&userNotificationPreferences); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingUserNotificationPreferences, err := c.userService.GetUserNotificationPreferencesByUserID(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Preferências de notificação não encontradas"})
		return
	}

	existingUserNotificationPreferences.Receberemails = userNotificationPreferences.Receberemails

	ctx.JSON(http.StatusOK, existingUserNotificationPreferences)
}

func Handler(router *gin.RouterGroup, userService *userServices.UserService, financeService *finanseServices.FinanceService) {
	controller := NewUserController(userService, financeService)

	router.POST("/users", controller.CreateUser)
	router.GET("/users/:id", controller.GetUserByID)
	router.PUT("/users/:id", controller.UpdateUser)
	router.DELETE("/users/:id", controller.DeleteUser)
	router.GET("/users", controller.ListUsers)
	router.POST("/users/notification-preferences", controller.CreateUserNotificationPreferences)
	router.PUT("/users/notification-preferences/:id", controller.UpdateUserNotificationPreferences)
	router.GET("/users/:id/balance", controller.GetUserBalance)
}
