package constants

const (
	ERR_API_INVALID_RISK_ID    = "Id is not a valid UUID"
	ERR_API_INVALID_RISK_STATE = "Risk state is invalid, must be one of ['open', 'closed', 'accepted', 'investigating']"
	ERR_API_PARSE_REQUEST_BODY = "Could not parse request body"
	ERR_API_RISK_NOT_FOUND     = "Risk not found"
)
