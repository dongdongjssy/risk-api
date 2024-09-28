package utils

var riskStates = [4]string{"open", "closed", "accepted", "investigating"}

// validate if given state is valid
func IsValidState(state string) bool {
	for _, val := range riskStates {
		if val == state {
			return true
		}
	}
	return false
}
