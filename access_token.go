package iiko

import (
	"github.com/google/uuid"
)

type AccessTokenRequest struct {
	// API login. It is set in iikoWeb [requried]
	ApiLogin string `json:"apiLogin"`
}

type AccessTokenResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// Authentication token. The standard token lifetime is 1 hour. [required]
	Token string `json:"token"`
}

// Retrieve session key for API user.
//
// iiko API: /api/1/access_token
func (c *Client) accessToken(req *AccessTokenRequest, opts ...Option) (*AccessTokenResponse, error) {
	var response AccessTokenResponse

	if err := c.post(false, "/api/1/access_token", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
