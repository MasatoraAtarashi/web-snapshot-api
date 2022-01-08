package ws

import "fmt"

type APIError struct {
	StatusCode int
	Message    string   `json:"message"`
	Errors     []string `json:"errors"`
}

var _ error = (*APIError)(nil)

func (err *APIError) Error() string {
	if err.Message == "" {
		return fmt.Sprintf("Error: StatusCode: %d, message: %s", err.StatusCode, err.Errors)
	}
	return fmt.Sprintf("Error: StatusCode: %d, message: %s", err.StatusCode, err.Message)
}
