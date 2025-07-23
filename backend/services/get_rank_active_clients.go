package services

import (
	"gestao-vendas/config"
	"gestao-vendas/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRankActiveClients(c *gin.Context) {
	var clients []models.Client
	if err := config.DB.Order("count_orders desc").Find(&clients).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar clientes"})
		return
	}

	var response []map[string]interface{}

	for _, client := range clients {
		rep := map[string]interface{}{
			"email":          client.Email,
			"createdAt":      client.CreatedAt,
			"ordersQuantity": client.CountOrders,
		}

		response = append(response, rep)
	}

	c.JSON(http.StatusOK, response)
}
