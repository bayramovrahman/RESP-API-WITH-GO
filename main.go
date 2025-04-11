package main

import (
	"example.com/rest-api/database"
	"example.com/rest-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	// 8080 portunda çalıştır
	server.Run(":8080")
}
