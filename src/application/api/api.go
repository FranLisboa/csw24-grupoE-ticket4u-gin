package api

import (
	eventController "const/application/api/controllers/event"
	tenantController "const/application/api/controllers/tenant"
	userController "const/application/api/controllers/user"
	event "const/core/services/event"
	tenant "const/core/services/tenant"
	user "const/core/services/user"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, db *sql.DB) {
	userService := user.NewService(db)
	eventService := event.NewEventoService(db)
	tenantService := tenant.NewTenantService(db)

	v1 := router.Group("/api/v1")

	userController.Handler(v1, userService)
	tenantController.Handler(v1, tenantService)
	eventController.Handler(v1, eventService)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

}
