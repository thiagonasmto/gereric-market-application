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
func CreateAdm(c *gin.Context) {
	var user models.Adm
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
func GetAdms(c *gin.Context) {
	var users []models.Adm
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

// READ ONE
func GetAdm(c *gin.Context) {
	id := c.Param("id")
	var user models.Adm
	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// UPDATE
func UpdateAdm(c *gin.Context) {
	id := c.Param("id")
	var user models.Adm

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	var input models.Adm
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	config.DB.Model(&user).Updates(input)
	c.JSON(http.StatusOK, user)
}

// DELETE
func DeleteAdm(c *gin.Context) {
	id := c.Param("id")
	var user models.Adm

	if err := config.DB.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao deletar"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Usuário deletado com sucesso"})
}
