package handlers

import (
	"github.com/walacysilvam/citassauro/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"

	"strconv"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// ROTA RAIZ
	// GET /
	r.GET("/quotes", func(c *gin.Context) {
		var quotes []models.Quote
		db.Find(&quotes)
		c.JSON(http.StatusOK, quotes)
	})

	// REGISTRO DE CITAÇÃO
	// POST /quotes
	r.POST("/quotes", func(c *gin.Context) {
		var quote models.Quote
		if err := c.ShouldBindJSON(&quote); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&quote)
		c.JSON(http.StatusCreated, quote)
	})

	// ATUALIZAÇÃO DE CITAÇÃO
	// PUT /quotes/:id
	r.PUT("/quotes/:id", func(c *gin.Context) {
		var quote models.Quote
		id := c.Param("id")

		//Se não existe 
		if err := db.First(&quote, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Citação não encontrada"})
			return
		} 

		var input models.Quote
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//update dps campos
		quote.Author = input.Author
		quote.Text = input.Text
		quote.Votes = input.Votes
		db.Save(&quote)
		c.JSON(http.StatusOK, quote)
	})

	// DELEÇÃO DE CITAÇÃO
	// DELETE /quotes/:id
	r.DELETE("/quotes/:id", func(c *gin.Context) {
		var quote models.Quote
		idParam := c.Param("id")
		// conversão de string para int
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
	})

	// REGISTRO DE USUARIOS
	// POST /users
	r.POST("/register", Register(db))
}