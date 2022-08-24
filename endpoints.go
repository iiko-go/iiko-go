package iiko

import "github.com/google/uuid"

// ErrorResponse is a custom error type that represents iikoCloud API Error.
type ErrorResponse struct {
	// http.StatusCode of request.
	StatusCode int `json:"-"`
	// Operation ID [required]
	CorrelationID uuid.UUID `json:"correlationId"`
	// Error text [required]
	ErrorDescription string `json:"errorDescription"`
	// Error code.
	ErrorField string `json:"error"`
}

// Error ...
func (e *ErrorResponse) Error() string {
	return e.ErrorDescription
}
