package api

import (
	"database/sql"
	hello-world....
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, db *sql.DB) {


	helloService := hello.NewService(db)


	v1 := router.Group("/api/v1")

	helloWorldController.Handler(v1, helloService)

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
}
