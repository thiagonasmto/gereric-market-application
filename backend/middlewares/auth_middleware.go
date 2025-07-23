package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("chave_de_assinatura")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Recupera o header Authorization: Bearer <token>
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			return
		}

		// Extrai o token do header
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Formato do token inválido"})
			return
		}

		tokenStr := parts[1]

		// Faz o parse e validação do token
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("algoritmo de assinatura inválido")
			}
			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token inválido ou expirado"})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			clientIDRaw, exists := claims["clientid"]
			if !exists {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Claim clientid não encontrada"})
				return
			}

			var clientID string

			switch v := clientIDRaw.(type) {
			case string:
				clientID = v
			case []byte:
				clientID = string(v)
			case float64:
				clientID = fmt.Sprintf("%.0f", v)
			default:
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Claim clientid tem tipo inválido"})
				return
			}

			c.Set("clientid", clientID)
			c.Next()
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Token malformado"})
	}
}
