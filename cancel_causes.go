package iiko

import (
	"github.com/google/uuid"
)

type CancelCauses struct {
	// Identifier. [required]
	ID uuid.UUID `json:"id"`

	// Name. [required]
	Name string `json:"name"`

	// Is deleted sign.
	IsDeleted bool `json:"isDeleted"`
}

type CancelCausesRequest struct {
	// Organizations ids which delivery cancel causes needs to be returned.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationIDs []uuid.UUID `json:"organizationIds"`
}

type CancelCausesResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// List of delivery cancel causes. [required]
	CancelCauses []CancelCauses `json:"cancelCauses"`
}

// Delivery cancel causes. Allowed from version 7.7.1.
//
// iiko API: /api/1/cancel_causes
func (c *Client) CancelCauses(req *CancelCausesRequest, opts ...Option) (*CancelCausesResponse, error) {
	var response CancelCausesResponse

	if err := c.post(true, "/api/1/cancel_causes", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
