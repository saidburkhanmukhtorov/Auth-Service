package api

import (
	"database/sql"

	_ "github.com/Project_Restaurant/Auth-Service/api/docs"
	"github.com/Project_Restaurant/Auth-Service/api/handler.go"
	"github.com/Project_Restaurant/Auth-Service/storage/postgres"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Auth service
// @version 1.0
// @description This is a Auth service.
// @host localhost:8081
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func NewRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	userRepo := postgres.NewUserRepo(db)
	h := handler.NewHandler(db, userRepo)
	router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", h.Register)
	router.POST("/login", h.Login)
	router.GET("/users/:username", h.GetByUsername)

	return router
}
