package handlers

import (
	"net/http"

	"github.com/dongdongjssy/risk-api/constants"
	"github.com/dongdongjssy/risk-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetRisks get a list of risks
// @Summary get a list of risks
// @Tags Risk
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /risks [get]
func GetRisks(ctx *gin.Context) {
	risks := models.GetRisks()
	ctx.JSON(http.StatusOK, risks)
}

// GetRisk get a risk by id
// @Summary get a risk by id
// @Tags Risk
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /risk [get]
func GetRisk(ctx *gin.Context) {
	idFromPath := ctx.Param("id")

	if _, err := uuid.Parse(idFromPath); err != nil {
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

// CreateRisk create a risk
// @Summary create a risk
// @Tags Risk
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /risks [post]
func CreateRisk(ctx *gin.Context) {
	var risk models.Risk

	if err := ctx.ShouldBindJSON(&risk); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": constants.ERR_API_PARSE_REQUEST_BODY,
		})
		return
	}

	if err := risk.Save(); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, risk)
}
