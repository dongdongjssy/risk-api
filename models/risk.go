package models

import (
	"errors"

	"github.com/dongdongjssy/risk-api/constants"
	"github.com/google/uuid"
)

// define risk structure
type Risk struct {
	ID          string `json:"id"`
	State       string `json:"state" validate:"required,oneof=open closed accepted investigating"`
	Title       string `json:"title" validate:"required,max=128"`
	Description string `json:"description" validate:"max=500"`
}

// use in-memory array to store all risks
var risks = []Risk{}

// save one risk into store
func (risk *Risk) Save() error {
	// validate title duplication
	for _, r := range risks {
		if r.Title == risk.Title {
			return errors.New(constants.ERR_MSG_DUPLICATE_RISK_TITLE)
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
