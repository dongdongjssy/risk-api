package models

type Risk struct {
	ID          string
	State       string `binding:"required"`
	Title       string `binding:"required"`
	Description string `binding:"required"`
}

// use in-memory array to store all risks
var risks = []Risk{}

// save one risk into store
func (risk Risk) Save() {
	risks = append(risks, risk)
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
