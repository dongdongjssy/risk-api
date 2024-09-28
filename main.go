package main

import (
	"log"

	"github.com/dongdongjssy/risk-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	// save logging to a local file that's in the same folder where the program runs
	// gin.DisableConsoleColor()
	// file, _ := os.Create("risk-api.log")
	// gin.DefaultWriter = io.MultiWriter(file)

	server := gin.Default()
	routes.RegisterRiskRoutes(server)
	server.Run("localhost:8080")
	log.Println("Risk api service is successfully running...")
}
