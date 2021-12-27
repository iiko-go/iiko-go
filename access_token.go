package iiko

import (
	"fmt"
)

type AccessTokenRequest struct {
	// API login. It is set in iikoWeb [requried]
	ApiLogin string `json:"apiLogin"`
}

type AccessTokenResponse struct {
	// Operation ID. [required]
	CorrelationID UUID `json:"correlationId"`

	// Authentication token. The standard token lifetime is 1 hour. [required]
	Token string `json:"token"`
}

// iiko API: /api/1/access_token
func (c *Client) AccessToken(req *AccessTokenRequest, opts ...Option) (*AccessTokenResponse, error) {
	var (
		onSuccess AccessTokenResponse
		onError   errorResponse
	)
	if err := c.post(false, "/api/1/access_token", req, &onSuccess, &onError, opts...); err != nil {
		return nil, err
	}
	if onError.ErrorDescription != "" {
		return nil, fmt.Errorf("iiko error: %q", onError.ErrorDescription)
	}
	return &onSuccess, nil
}
