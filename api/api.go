package api

import (
	"database/sql"

	"github.com/Project_Restaurant/Auth-Service/api/handler.go"
	"github.com/gin-gonic/gin"
)

func NewGinRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()
	handler := handler.NewHandler(db)

	router.POST("/register", handler.Register)
	router.POST("/login", handler.Login)
	router.GET("/user/:username", handler.GetByUsername)

	return router
}
