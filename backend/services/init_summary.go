package services

import (
	"gestao-vendas/config"
	"gestao-vendas/models"
)

func InitSalesSummaryIfNotExists() {
	var count int64
	config.DB.Model(&models.SalesSummary{}).Count(&count)

	if count == 0 {
		config.DB.Create(&models.SalesSummary{
			ID:              0,
			TotalSalesMade:  0,
			InvoicedAmount:  0.0,
			TotalOrdersSold: 0,
		})
	}
}
