package handlers

import (
	"github.com/walacysilvam/citassauro/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/quotes", func(c *gin.Context) {
		var quotes []models.Quote
		db.Find(&quotes)
		c.JSON(http.StatusOK, quotes)
	})

	r.POST("/quotes", func(c *gin.Context) {
		var quote models.Quote
		if err := c.ShouldBindJSON(&quote); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&quote)
		c.JSON(http.StatusCreated, quote)
	})

}