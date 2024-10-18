package api

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine, db *sql.DB) {
	router.Group("/api/v1")

	// v1.GET("/users", func(ctx *gin.Context) {
	// 	getUsers(ctx, db)
	// })

	// v1.GET("/users/:id", func(ctx *gin.Context) {
	// 	getUserByID(ctx, db)
	// })

	// v1.POST("/users", func(ctx *gin.Context) {
	// 	createUser(ctx, db)
	// })

	// v1.PUT("/users/:id", func(ctx *gin.Context) {
	// 	updateUser(ctx, db)
	// })

	// v1.DELETE("/users/:id", func(ctx *gin.Context) {
	// 	deleteUser(ctx, db)
	// })
}
