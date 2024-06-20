package middlewares

import (
	"fmt"
	"log"
	"net/http"

	tokens "github.com/Project_Restaurant/Auth-Service/api/token"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth(c *gin.Context) string {
	tokenString := c.GetHeader("Authorization")
	log.Println(tokenString)
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
		c.Abort()
		return ""
	}
	if len(tokenString) < 8 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong authorization header"})
		c.Abort()
		return ""
	}
	tokenString = tokenString[len("Bearer "):]
	err := tokens.VerifyToken(tokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return ""
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return ""
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		c.Abort()
		return ""
	}
	c.Next()
	return fmt.Sprint(claims["username"])
}
