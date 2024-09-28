package routes

import (
	"github.com/dongdongjssy/risk-api/handlers"
	"github.com/dongdongjssy/risk-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRiskRoutes(server *gin.Engine) {
	riskApis := server.Group("/")

	// assume the apis can be used by only valid users or services
	riskApis.Use(middlewares.Authenticate)
	riskApis.GET("/v1/risks", handlers.GetRisks)
	riskApis.GET("/v1/risks/:id", handlers.GetRisk)
	riskApis.POST("/v1/risks", handlers.CreateRisk)
}
