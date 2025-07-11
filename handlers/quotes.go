package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/walacysilvam/citassauro/models"
	"gorm.io/gorm"
)

func GetQuotes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var quotes []models.Quote
		db.Find(&quotes)
		c.JSON(http.StatusOK, quotes)
	}
}

func CreateQuote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var quote models.Quote
		if err := c.ShouldBindJSON(&quote); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&quote)
		c.JSON(http.StatusCreated, quote)
	}
}

func UpdateQuote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var quote models.Quote
		id := c.Param("id")

		if err := db.First(&quote, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Citação não encontrada"})
			return
		}

		var input models.Quote
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		quote.Author = input.Author
		quote.Text = input.Text
		quote.Votes = input.Votes
		db.Save(&quote)
		c.JSON(http.StatusOK, quote)
	}
}

func DeleteQuote(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var quote models.Quote
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		if err := db.First(&quote, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Citação não encontrada"})
			return
		}

		db.Delete(&quote)
		c.Status(http.StatusNoContent)
	}
}