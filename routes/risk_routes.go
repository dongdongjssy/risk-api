package routes

import (
	"github.com/dongdongjssy/risk-api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRiskRoutes(server *gin.Engine) {
	server.GET("/v1/risks", handlers.GetRisks)
	server.GET("/v1/risks/:id", handlers.GetRisk)
	server.POST("/v1/risks", handlers.CreateRisk)
}
