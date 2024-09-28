package handlers

import (
	"net/http"

	"github.com/dongdongjssy/risk-api/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetRisks(ctx *gin.Context) {
	risks := models.GetRisks()
	ctx.JSON(http.StatusOK, risks)
}

func GetRisk(ctx *gin.Context) {
	_, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse risk id."})
		return
	}
	ctx.JSON(http.StatusOK, nil)
}

func CreateRisk(ctx *gin.Context) {
	var risk models.Risk

	err := ctx.ShouldBindJSON(&risk)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request body"})
		return
	}

	risk.ID = uuid.New()
	risk.Save()

	ctx.JSON(http.StatusCreated, risk)
}
