package token

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var SecurityKey = []byte("This is true")

func CreateToken(id, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": username,
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Add(10 * time.Minute).Unix(),
	})

	tokenSecredCode, err := token.SignedString(SecurityKey)
	if err != nil {
		log.Println("Error creating token:", err)
		return "", err
	}
	return tokenSecredCode, nil
}

func ProtectedToken(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Missing authorization header"})
		return
	}
	token = token[len("Bearer "):]
	err := verifyToken(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid token"})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Welcome to the protected area"})
}

func verifyToken(tokenreq string) error {
	token, err := jwt.Parse(tokenreq, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return SecurityKey, nil
	})
	if err != nil {
		log.Println("Error parsing token:", err)
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}
