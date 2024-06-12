package handler

import (
	"auth/models"
	"auth/token"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) RegisterUser(c *gin.Context)  {
	userReq := models.UserReq{}
	err := c.BindJSON(&userReq)
	if err != nil{
		log.Fatal("Error while bindjson in registeruser",err)
		return
	}
	res,err := h.UserDb.CreateUser(&userReq)
	if err != nil{
		c.JSON(400,err)
		return
	}
	token,err := token.CreateToken(res.ID,res.Name)
	if err != nil{
		log.Fatal("Error while Token in registeruser",err)
		return
	}
	res.Token = token
	c.JSON(http.StatusOK,res)
}
func (h Handler) LoginUser(c *gin.Context)  {
	
}
func (h Handler) UpdaterUser(c *gin.Context)  {
	
}
func (h Handler) DeleteUser(c *gin.Context)  {
	
}
func (h Handler) GetUser(c *gin.Context)  {
	
}

