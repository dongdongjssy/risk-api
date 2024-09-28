package models

type Risk struct {
	ID          [16]byte
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

func GetRiskById(id [16]byte) (*Risk, error) {
	for _, risk := range risks {
		if risk.ID == id {
			return &risk, nil
		}
	}

	return nil, nil
}
