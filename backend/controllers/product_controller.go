package controllers

import (
	"gestao-vendas/config"
	"gestao-vendas/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

// CREATE
func CreateProduct(c *gin.Context) {
	var product models.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar UUID para produto"})
		return
	}
	product.ProductID = uid

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	if err := config.DB.Create(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar produto"})
		return
	}
	c.JSON(http.StatusCreated, product)
}

// READ ALL
func GetProducts(c *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	c.JSON(http.StatusOK, products)
}

// READ ONE
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// UPDATE
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&product).Updates(input)
	c.JSON(http.StatusOK, product)
}

// DELETE
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if err := config.DB.Where("id = ?", id).First(&product).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Produto não encontrado"})
		return
	}

	if err := config.DB.Delete(&product).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao deletar produto"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Produto deletado com sucesso"})
}
