package iiko

import (
	"github.com/google/uuid"
)

type RemovalType struct {
	// Identifier. [required]
	ID uuid.UUID `json:"id"`

	// Name of removal type. [required]
	Name string `json:"name"`

	// Comment.
	Comment string `json:"comment"`

	// Can write off to cafe.
	CanWriteoffToCafe bool `json:"canWriteoffToCafe"`

	// Can write off to waiter.
	CanWriteoffToWaiter bool `json:"canWriteoffToWaiter"`

	// Can write off to user.
	CanWriteoffToUser bool `json:"canWriteoffToUser"`

	// Require comments on operations.
	ReasonRequired bool `json:"reasonRequired"`

	// Can be used manually.
	Manual bool `json:"manual"`

	// Is deleted sign.
	IsDeleted bool `json:"isDeleted"`
}

type RemovalTypesRequest struct {
	// Organizations ids which removal types needs to be returned.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationIDs []uuid.UUID `json:"organizationIds"`
}

type RemovalTypesResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// List of removal types. [required]
	RemovalTypes []RemovalType `json:"paymentTypes"`
}

// Removal types (reasons for deletion). Allowed from version 7.5.3.
//
// iiko API: /api/1/removal_types
func (c *Client) RemovalTypes(req *RemovalTypesRequest, opts ...Option) (*RemovalTypesResponse, error) {
	var response RemovalTypesResponse

	if err := c.post(true, "/api/1/removal_types", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
