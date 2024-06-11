package handler

import (
	"database/sql"
	"net/http"

	"github.com/Project_Restaurant/Auth-Service/models"
	"github.com/Project_Restaurant/Auth-Service/postgres"
	"github.com/Project_Restaurant/Auth-Service/token"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	db   *sql.DB
	User *postgres.UserRepo
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		db:   db,
		User: postgres.NewUserRepo(db),
	}
}

func (h *Handler) Register(c *gin.Context) {
	var user models.UserRegister
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := h.User.Register(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := token.CreateToken(res.Name, res.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Login(c *gin.Context) {
	var user models.UserLogin
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.User.Login(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, err := token.CreateToken(res.Name, res.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) GetByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := h.User.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

