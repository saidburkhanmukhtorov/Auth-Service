package api

import (
	"auth/api/handler"
	"auth/postgres"
	"database/sql"

	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

func NewEngine(db *sql.DB) *gin.Engine {
	router := gin.Default()
	user := postgres.NewUserRepo(db)
	h := handler.NewHandler(db, user)
	// router.Use(token.ProtectedToken)
	// router.GET("api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/register", h.RegisterUser)
	router.POST("/login", h.LoginUser)
	router.POST("/get/:id", h.GetUser)
	router.PUT("/update", h.UpdaterUser)
	router.DELETE("/delete", h.DeleteUser)
	return router
}
