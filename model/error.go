package model

import "fmt"

// RequestError struct
type RequestError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewRequestError function
func NewRequestError(code string, message string) error {
	return &RequestError{
		Code:    code,
		Message: message,
	}
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("Code: %s\nMessage:%s", r.Code, r.Message)
}
