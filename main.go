package main

import (
	"github.com/dongdongjssy/risk-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	routes.RegisterRiskRoutes(server)
	server.Run("localhost:8080")
}
