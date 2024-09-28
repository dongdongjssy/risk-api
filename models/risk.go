package models

type Risk struct {
	ID          string
	State       string `binding:"required"`
	Title       string `binding:"required"`
	Description string `binding:"required"`
}

var risks = []Risk{}

func (risk Risk) Save() {
	risks = append(risks, risk)
}

func GetRisks() []Risk {
	return risks
}

func GetRiskById(id string) (*Risk, error) {
	for _, risk := range risks {
		if risk.ID == id {
			return &risk, nil
		}
	}

	return nil, nil
}
