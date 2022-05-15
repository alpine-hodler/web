package client

import (
	"encoding/json"
)

type ErrorMessage struct {
	Message    string `json:"message"`
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
	Detail     string `json:"detail"`
	Title      string `json:"title"`
	Type       string `json:"type"`
}

func (msg *ErrorMessage) UnmarshalJSON(b []byte) error {
	var data map[string]interface{}
	if err := json.Unmarshal(b, &data); err != nil {
		return err
	}

	if val := data["message"]; val != nil {
		msg.Message = val.(string)
	}
	if val := data["detail"]; val != nil {
		msg.Detail = val.(string)
	}
	if val := data["title"]; val != nil {
		msg.Title = val.(string)
	}
	if val := data["type"]; val != nil {
		msg.Type = val.(string)
	}

	return nil
}
