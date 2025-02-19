package errors

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"` // Optional, for additional error information
}
