package utils

type ApplicationError struct {
	Message     string `json:"error_message"`
	StatusCode  int    `json:"status_code"`
	Description string `json:"description"`
}
