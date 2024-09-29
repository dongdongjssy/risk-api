package main

import (
	"github.com/dongdongjssy/risk-api/routes"
	"github.com/gin-gonic/gin"
)

// @title Risk APIs
// @version 1.0
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /v1
func main() {
	server := gin.Default()
	routes.RegisterRiskRoutes(server)
	server.Run("localhost:8080")
}
