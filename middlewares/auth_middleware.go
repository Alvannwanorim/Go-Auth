package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
	}
	token, err := ValidateToken(tokenString)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
	}
	claims := token.Claims.(jwt.MapClaims)

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
	}

	email, ok := claims["email"]
	if !ok || email == "" {
		c.AbortWithError(http.StatusUnauthorized, errors.New("Unauthorized"))
	}
	c.Set("user", email)

	c.Next()
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	JWT_SECRET := os.Getenv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected sign method %v", token.Header["alg"])
		}
		return []byte(JWT_SECRET), nil
	})
}
