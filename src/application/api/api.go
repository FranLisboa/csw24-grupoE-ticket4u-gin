package api

import (
	userController "const/application/api/controllers/user"
	"const/core/services/user"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Setup(router *gin.Engine, db *sql.DB) {
	userService := user.NewService(db)

	v1 := router.Group("/api/v1")

	userController.Handler(v1, userService)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
