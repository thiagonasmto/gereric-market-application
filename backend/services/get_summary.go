package services

import (
	"gestao-vendas/config"
	"gestao-vendas/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSummary(c *gin.Context) {
	var summary models.SalesSummary
	if err := config.DB.Find(&summary).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao consultar sumário de vendas"})
		return
	}
	c.JSON(http.StatusOK, summary)
}
