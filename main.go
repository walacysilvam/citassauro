package main

import (
	"github.com/gin-gonic/gin"
	"github.com/walacysilvam/citassauro/handlers"
	"github.com/walacysilvam/citassauro/utils"
	"github.com/walacysilvam/citassauro/models"
)

func main() {
	db := utils.InitDB()
	r := gin.Default()

	models.AutoMigrate(db)

	handlers.RegisterRoutes(r, db)

	r.Run(":8080")
}
