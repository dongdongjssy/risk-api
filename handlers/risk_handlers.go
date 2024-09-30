package handlers

import (
	"log"
	"net/http"

	"github.com/dongdongjssy/risk-api/constants"
	"github.com/dongdongjssy/risk-api/models"
	"github.com/dongdongjssy/risk-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

// GetRisks get a list of risks
//
// @Summary get a list of risks
// @Tags Risk
// @Accept json
// @Produce json
// @Success 200 {array} models.Risk
// @Success 500 {string} string "internal server error"
// @Router /risks [get]
func GetRisks(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}()

	risks := models.GetRisks()
	ctx.JSON(http.StatusOK, risks)
}

// GetRisk get a risk by id
//
// @Summary get a risk by id
// @Tags Risk
// @Accept json
// @Produce json
// @Param id query string true "risk uuid"
// @Success 200 {object} models.Risk
// @Failure	400 {string} string "bad request"
// @Failure	404 {string} string "not found"
// @Success 500 {string} string "internal server error"
// @Router /risk/{id} [get]
func GetRisk(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}()

	idFromPath := ctx.Param("id")

	// validate uuid from request
	if _, err := uuid.Parse(idFromPath); err != nil {
		ctx.JSON(http.StatusBadRequest, []string{constants.ERR_API_INVALID_RISK_ID})
		log.Println(utils.FormatLog("GetRisk", err.Error()))
		return
	}

	// search from store
	risk := models.GetRiskById(idFromPath)
	if risk == nil {
		ctx.JSON(http.StatusNotFound, []string{constants.ERR_API_RISK_NOT_FOUND})
		log.Println(utils.FormatLog("GetRisk", constants.ERR_API_RISK_NOT_FOUND))
		return
	}

	ctx.JSON(http.StatusOK, risk)
}

// CreateRisk create a risk
//
// @Summary create a risk
// @Tags Risk
// @Accept json
// @Produce json
// @Param risk body models.Risk true "a new risk (without id)"
// @Success 200 {object} models.Risk "new created risk with uuid"
// @Failure	400 {string} string "bad request"
// @Success 500 {string} string "internal server error"
// @Router /risks [post]
func CreateRisk(ctx *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
	}()

	var risk models.Risk

	// parse request body and do validation
	validate := validator.New()
	if err := validate.Struct(risk); err != nil {
		errMsgs := utils.ParseValidationErr(&err)
		ctx.JSON(http.StatusBadRequest, errMsgs)
		log.Println(utils.FormatLog("CreateRisk", errMsgs))
		return
	}

	// save risk to store
	if err := risk.Save(); err != nil {
		ctx.JSON(http.StatusBadRequest, []string{err.Error()})
		log.Println(utils.FormatLog("CreateRisk", err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, risk)
}
