package services

import (
	"gestao-vendas/config"
	"gestao-vendas/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrdersInProgress(c *gin.Context) {
	var orders []models.Order

	if err := config.DB.Where("status = ? AND processed = 0", "Em andamento").Preload("Client").Preload("Products.Product").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar pedidos em andamento"})
		return
	}

	var response []map[string]interface{}

	for _, order := range orders {
		var products []map[string]interface{}
		for _, p := range order.Products {
			products = append(products, map[string]interface{}{
				"id":       p.ProductID,
				"quantity": p.Quantity,
				"name":     p.Product.Name,
			})
		}

		rep := map[string]interface{}{
			"id": order.ID,
			"client": map[string]interface{}{
				"id":    order.Client.ID,
				"email": order.Client.Email,
			},
			"products":   products,
			"totalPrice": order.TotalPrice,
			"status":     order.Status,
		}

		response = append(response, rep)
	}

	c.JSON(http.StatusOK, response)
}
