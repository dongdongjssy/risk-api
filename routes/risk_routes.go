package routes

import (
	"github.com/dongdongjssy/risk-api/constants"
	"github.com/dongdongjssy/risk-api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRiskRoutes(server *gin.Engine) {
	server.GET(constants.ENDPOINT_GET_RISKS, handlers.GetRisks)
	server.GET(constants.ENDPOINT_GET_RISK, handlers.GetRisk)
	server.POST(constants.ENDPOINT_POST_RISKS, handlers.CreateRisk)
}
