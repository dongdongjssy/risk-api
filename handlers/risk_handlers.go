package handlers

import (
	"net/http"
	"strconv"

	"github.com/dongdongjssy/risk-api/models"
	"github.com/gin-gonic/gin"
)

func GetRisks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, []models.Risk{})
}

func GetRisk(ctx *gin.Context) {
	_, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
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

	ctx.JSON(http.StatusCreated, gin.H{"message": "created"})
}
