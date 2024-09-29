package routes

import (
	"github.com/dongdongjssy/risk-api/constants"
	_ "github.com/dongdongjssy/risk-api/docs"
	"github.com/dongdongjssy/risk-api/handlers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRiskRoutes(server *gin.Engine) {
	// register swagger ui path
	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// register risk api endpoints
	server.GET(constants.ENDPOINT_GET_RISKS, handlers.GetRisks)
	server.GET(constants.ENDPOINT_GET_RISK, handlers.GetRisk)
	server.POST(constants.ENDPOINT_POST_RISKS, handlers.CreateRisk)
}
