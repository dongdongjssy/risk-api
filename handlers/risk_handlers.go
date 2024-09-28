package handlers

import (
	"net/http"

	"github.com/dongdongjssy/risk-api/constants"
	"github.com/dongdongjssy/risk-api/models"
	"github.com/dongdongjssy/risk-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRisks(ctx *gin.Context) {
	risks := models.GetRisks()
	ctx.JSON(http.StatusOK, risks)
}

func GetRisk(ctx *gin.Context) {
	idFromPath := ctx.Param("id")

	_, err := uuid.Parse(idFromPath)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.ERR_API_INVALID_RISK_ID,
		})
		return
	}

	risk := models.GetRiskById(idFromPath)
	if risk == nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": constants.ERR_API_RISK_NOT_FOUND,
		})
		return
	}

	ctx.JSON(http.StatusOK, risk)
}

func CreateRisk(ctx *gin.Context) {
	var risk models.Risk

	err := ctx.ShouldBindJSON(&risk)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.ERR_API_PARSE_REQUEST_BODY,
		})
		return
	}

	if isValidState := utils.IsValidState(risk.State); !isValidState {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.ERR_API_INVALID_RISK_STATE,
		})
		return
	}

	risk.ID = uuid.New().String()
	risk.Save()

	ctx.JSON(http.StatusCreated, risk)
}
