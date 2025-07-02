package utils

import (
	"github.com/walacysilvam/citassauro/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {

	db, err := gorm.Open(sqlite.Open("citassauro.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	db.AutoMigrate(&models.Quote{})
	return db
}
