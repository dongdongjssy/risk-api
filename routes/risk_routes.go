package routes

import (
	"github.com/dongdongjssy/risk-api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRiskRoutes(server *gin.Engine) {
	authenticated := server.Group("/")

	// assume the apis are used for valid users or services
	authenticated.Use(nil)
	authenticated.GET("/v1/risks", handlers.GetRisks)
	authenticated.GET("/v1/risks/:id", handlers.GetRisk)
	authenticated.POST("/v1/risks", handlers.CreateRisk)
}
