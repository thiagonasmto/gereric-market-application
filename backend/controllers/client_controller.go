package controllers

import (
	"gestao-vendas/config"
	"gestao-vendas/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// CREATE
func CreateClient(c *gin.Context) {
	var user models.Client
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := uuid.NewV4()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar UUID"})
		return
	}
	user.ID = uid

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criptografar senha"})
		return
	}
	user.Password = string(hashedPassword)

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// READ ALL
func GetClients(c *gin.Context) {
	var users []models.Client
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// READ ONE
func GetClient(c *gin.Context) {
	id := c.Param("id")
	var user models.Client
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UPDATE
func UpdateClient(c *gin.Context) {
	id := c.Param("id")
	var user models.Client

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	var input models.Client
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, user)
}

// DELETE
func DeleteClient(c *gin.Context) {
	id := c.Param("id")
	var user models.Client

	// Primeiro verifica se o cliente existe
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}
