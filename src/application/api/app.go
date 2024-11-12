package api

import (
    "database/sql"
    "github.com/gin-gonic/gin"
)

type App struct {
    DB     *sql.DB
    Router *gin.Engine
}

func NewApp(db *sql.DB) *App {
    app := &App{
        DB:     db,
        Router: gin.Default(),
    }
    Setup(app.Router, app.DB)
    return app
}


//teste