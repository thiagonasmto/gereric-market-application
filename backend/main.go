package main

import (
	"fmt"
	"gestao-vendas/config"
	"gestao-vendas/models"
	"gestao-vendas/routes"
	"gestao-vendas/services"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	// conectar ao banco
	config.Connect()

	fmt.Println("Migrando tabelas...")
	if err := config.DB.AutoMigrate(&models.Client{}); err != nil {
		log.Fatalf("Erro ao migrar UserModel: %v", err)
	}
	if err := config.DB.AutoMigrate(&models.Adm{}); err != nil {
		log.Fatalf("Erro ao migrar AdmModel: %v", err)
	}
	if err := config.DB.AutoMigrate(&models.Product{}); err != nil {
		log.Fatalf("Erro ao migrar ProductModel: %v", err)
	}
	if err := config.DB.AutoMigrate(&models.OrderProduct{}); err != nil {
		log.Fatalf("Erro ao migrar OrderProduct: %v", err)
	}
	if err := config.DB.AutoMigrate(&models.Order{}); err != nil {
		log.Fatalf("Erro ao migrar OrderModel: %v", err)
	}
	if err := config.DB.AutoMigrate(&models.SalesSummary{}); err != nil {
		log.Fatalf("Erro ao migrar SalesSummaryModel: %v", err)
	} else {
		services.InitSalesSummaryIfNotExists()
	}

	go func() {
		for {
			var orders []models.Order
			if err := config.DB.Where("status = ? AND processed = 0", "Finalizado").Preload("Products").Find(&orders).Error; err == nil {
				var summary models.SalesSummary
				if err := config.DB.First(&summary).Error; err != nil {
					log.Println("Resumo de vendas não encontrado:", err)
					time.Sleep(10 * time.Second)
					continue
				}

				totalOrdersSold := 0
				totalSalesMade := 0
				totalInvoiced := 0.0

				for _, order := range orders {
					totalOrdersSold++

					for _, produto := range order.Products {
						var p models.Product
						if err := config.DB.First(&p, "product_id = ?", produto.ProductID).Error; err != nil {
							log.Println("Produto não encontrado:", err)
							continue
						}

						if p.Quantity < produto.Quantity {
							log.Println("Estoque insuficiente para produto:", produto.ProductID)
							continue
						}

						if err := config.DB.Model(&models.Product{}).
							Where("product_id = ?", produto.ProductID).
							Update("quantity", gorm.Expr("quantity - ?", produto.Quantity)).Error; err != nil {
							log.Println("Erro ao atualizar estoque:", err)
							continue
						}

						totalSalesMade += produto.Quantity
					}

					totalInvoiced += order.TotalPrice

					// Atualiza pedido como processado
					config.DB.Model(&order).Update("processed", 1)

					var client models.Client
					if err := config.DB.Model(&client).
						Where("id = ?", order.ClientID).
						Update("count_orders", gorm.Expr("count_orders + ?", 1)).Error; err != nil {
						log.Println("Erro ao atualizar quantidade de pedidos realizados pelo cliente:", err)
						continue
					}
				}

				config.DB.Model(&summary).Where("id = ?", 0).Updates(models.SalesSummary{
					TotalOrdersSold: summary.TotalOrdersSold + totalOrdersSold,
					TotalSalesMade:  summary.TotalSalesMade + totalSalesMade,
					InvoicedAmount:  summary.InvoicedAmount + totalInvoiced,
				})
			} else {
				log.Println("Erro ao buscar pedidos:", err)
			}

			time.Sleep(10 * time.Second)
		}
	}()

	// iniciar servidor
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// rotas
	routes.UserRoutes(r)

	// iniciar na porta 8081
	r.Run(":8081")
}
