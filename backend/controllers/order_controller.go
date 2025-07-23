package controllers

import (
	"fmt"
	"gestao-vendas/config"
	"gestao-vendas/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func CreateOrder(c *gin.Context) {
	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, p := range order.Products {
		var product models.Product
		if err := config.DB.Select("quantity").Where("product_id = ?", p.ProductID).First(&product).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Produto %s não encontrado", p.ProductID)})
			return
		}

		if p.Quantity > product.Quantity {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Estoque insuficiente para produto %s", p.ProductID)})
			return
		}
	}

	uid, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar UUID para pedido"})
		return
	}
	order.ID = uid

	totalPrice := 0.0
	for _, produto := range order.Products {
		var prod models.Product
		if err := config.DB.Select("price").Where("product_id = ?", produto.ProductID).First(&prod).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar preço do produto"})
			return
		}
		totalPrice += prod.Price * float64(produto.Quantity)
	}

	order.TotalPrice = totalPrice

	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()

	order.Status = "Em andamento"

	if err := config.DB.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar pedido"})
		return
	}

	if err := config.DB.Preload("Client").First(&order, "id = ?", order.ID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao carregar dados do pedido"})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order

	err := config.DB.
		Preload("Client").
		Preload("Products.Product").
		Find(&orders).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pedidos"})
		return
	}

	var response []map[string]interface{}
	for _, order := range orders {
		var products []map[string]interface{}
		for _, produto := range order.Products {
			products = append(products, map[string]interface{}{
				"id":       produto.ProductID,
				"quantity": produto.Quantity,
				"name":     produto.Product.Name,
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

	var string_response string
	if response == nil {
		string_response = "Nenhum pedido cadastrado"
		c.JSON(http.StatusOK, string_response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func GetOrderById(c *gin.Context) {
	id := c.Param("id")
	var order models.Order
	if err := config.DB.Where("id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido não encontrado"})
		return
	}

	err := config.DB.
		Preload("Client").
		Preload("Products").
		Find(&order).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pedidos"})
		return
	}

	var response []map[string]interface{}
	var products []map[string]interface{}
	var product models.Product
	for _, produto := range order.Products {
		if err := config.DB.Model(&product).Select("name").Where("id = ?", produto.ProductID).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produto"})
			return
		}
		products = append(products, map[string]interface{}{
			"id":       produto.ProductID,
			"quantity": produto.Quantity,
			"name":     product.Name,
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

	c.JSON(http.StatusOK, response)
}

func UpdateOrder(c *gin.Context) {
	id := c.Param("id")
	var order models.Order

	if err := config.DB.Where("id = ?", id).First(&order).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pedido não encontrado"})
		return
	}

	err := config.DB.
		Preload("Client").
		Preload("Products").
		Find(&order).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar pedidos"})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&order).Updates(input)

	var response []map[string]interface{}
	var products []map[string]interface{}
	for _, p := range order.Products {
		products = append(products, map[string]interface{}{
			"id":       p.ProductID,
			"quantity": p.Quantity,
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

	c.JSON(http.StatusOK, response)
}
