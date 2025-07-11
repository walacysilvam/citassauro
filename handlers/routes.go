package handlers

import (

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	// ROTAS DE QUOTES(citações)
	r.GET("/quotes", GetQuotes(db))
	r.POST("/quotes", CreateQuote(db))
	r.PUT("/quotes/:id", UpdateQuote(db))
	r.DELETE("/quotes/:id", DeleteQuote(db))

	// ROTAS DE USUÁRIOS
	r.POST("/register", Register(db))
	r.POST("/login", Login(db))
	r.DELETE("/users/:id", deleteUser(db))
}