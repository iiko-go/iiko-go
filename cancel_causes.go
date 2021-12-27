package iiko

import "fmt"

type CancelCauses struct {
	// Identifier. [required]
	ID UUID `json:"id"`

	// Name. [required]
	Name string `json:"name"`

	// Is deleted sign.
	IsDeleted bool `json:"isDeleted"`
}

type CancelCausesRequest struct {
	// Organizations ids which delivery cancel causes needs to be returned.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationIDs []UUID `json:"organizationIds"`
}

type CancelCausesResponse struct {
	// Operation ID. [required]
	CorrelationID UUID `json:"correlationId"`

	// List of delivery cancel causes. [required]
	CancelCauses []CancelCauses `json:"cancelCauses"`
}

// iiko API: /api/1/cancel_causes
func (c *Client) CancelCauses(req *CancelCausesRequest, opts ...Option) (*CancelCausesResponse, error) {
	var (
		onSuccess CancelCausesResponse
		onError   errorResponse
	)
	if err := c.post(true, "/api/1/cancel_causes", req, &onSuccess, &onError, opts...); err != nil {
		return nil, err
	}
	if onError.ErrorDescription != "" {
		return nil, fmt.Errorf("iiko error: %q", onError.ErrorDescription)
	}
	return &onSuccess, nil
}
