package models

type Risk struct {
	ID          [16]byte
	State       string `binding:"required"`
	Title       string `binding:"required"`
	Description string `binding:"required"`
}

var risks = []Risk{}

func (r Risk) Save() {
	risks = append(risks, r)
}

func GetRisks() []Risk {
	return risks
}

func GetRiskById(id int64) (*Risk, error) {
	return nil, nil
}
