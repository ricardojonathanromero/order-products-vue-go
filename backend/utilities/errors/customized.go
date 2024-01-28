package errors

import "encoding/json"

type CustomError struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details"`
}

func (c *CustomError) Error() string {
	m, _ := json.Marshal(c)
	return string(m)
}
