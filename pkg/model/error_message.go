package model

type ErrorMessage struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
}
