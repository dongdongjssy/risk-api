package models

type ErrorMessage struct {
	Code    string   `json:"code"`
	Message []string `json:"message"`
}
