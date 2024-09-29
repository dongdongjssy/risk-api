package constants

const (
	// Error api responses
	ERR_API_INVALID_RISK_ID      = "id is not a valid UUID"
	ERR_API_INVALID_RISK_STATE   = "risk state is invalid, must be one of ['open', 'closed', 'accepted', 'investigating']"
	ERR_API_PARSE_REQUEST_BODY   = "could not parse request body"
	ERR_API_RISK_NOT_FOUND       = "risk not found"
	ERR_API_RISK_DUPLICATE_TITLE = "duplicate risk title, please choose another title"
)
