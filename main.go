package main

import (
	"github.com/gin-gonic/gin"
	"github.com/walacysilvam/citassauro/handlers"
	"github.com/walacysilvam/citassauro/utils"
)

func main() {
	db := utils.InitDB()
	r := gin.Default()

	handlers.RegisterRoutes(r, db)

	r.Run(":8080")
}
