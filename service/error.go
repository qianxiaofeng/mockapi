package service

import "fmt"

type SaveNodeError struct {
	ErrorCode int    `json:"error_code"`
	Message   string `json:"message"`
}

func (e SaveNodeError) Error() string {
	return fmt.Sprintf("SaveNodeError : %s", e.Message)
}
