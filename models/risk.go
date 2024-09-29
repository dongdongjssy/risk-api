package models

import (
	"errors"

	"github.com/dongdongjssy/risk-api/constants"
	"github.com/dongdongjssy/risk-api/utils"
	"github.com/google/uuid"
)

// define risk structure
type Risk struct {
	ID          string `json:"id"`
	State       string `json:"state" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

// use in-memory array to store all risks
var risks = []Risk{}

// save one risk into store
func (risk *Risk) Save() error {
	// validate state
	if isValidState := utils.IsValidState(risk.State); !isValidState {
		return errors.New(constants.ERR_API_INVALID_RISK_STATE)
	}

	// validate title duplication
	for _, r := range risks {
		if r.Title == risk.Title {
			return errors.New(constants.ERR_API_RISK_DUPLICATE_TITLE)
		}
	}

	// generate a random uuid and save it to array
	risk.ID = uuid.New().String()
	risks = append(risks, *risk)
	return nil
}

// get a list of risk from store
func GetRisks() []Risk {
	return risks
}

// find a risk by id in the store
func GetRiskById(id string) *Risk {
	for _, risk := range risks {
		if risk.ID == id {
			return &risk
		}
	}

	return nil
}
