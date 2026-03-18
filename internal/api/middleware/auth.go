package middleware

import (
	"fmt"
	"maps"
	"net/http"
	"strings"

	"github.com/BramAristyo/go-pos-mawish/internal/infrastructure/config"
	"github.com/BramAristyo/go-pos-mawish/pkg/service_errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		token := strings.Split(auth, " ")

		if auth == "" || len(token) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token required"})
			return
		}

		at, err := jwt.Parse(token[1], func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, service_errors.TokenInvalid
			}
			return []byte(cfg.JWT.Secret), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		claims, ok := at.Claims.(jwt.MapClaims)
		if !ok || !at.Valid {
			fmt.Println("Failed")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		claimMap := make(map[string]any)
		maps.Copy(claimMap, claims)

		c.Set("userID", claimMap["userID"])
		c.Set("role", claimMap["role"])
		c.Next()
	}
}
