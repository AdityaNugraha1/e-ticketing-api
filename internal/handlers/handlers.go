package handlers

import (
	"e-ticketing-api/internal/auth"
	"e-ticketing-api/internal/database"
	"e-ticketing-api/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var admin models.Admin
	db := database.DB
	if err := db.Where("username = ?", input.Username).First(&admin).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	if !auth.CheckPasswordHash(input.Password, admin.PasswordHash) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Username atau password salah"})
		return
	}

	token, err := auth.GenerateJWT(admin.Username, admin.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal membuat token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

type CreateTerminalInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateTerminalHandler(c *gin.Context) {
	var input CreateTerminalInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	role, exists := c.Get("role")
	if !exists || role.(string) != "superadmin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Akses ditolak. Hanya superadmin yang dapat membuat terminal."})
		return
	}

	terminal := models.Terminal{Name: input.Name}
	db := database.DB
	if err := db.Create(&terminal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal menyimpan terminal ke database"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": terminal})
}