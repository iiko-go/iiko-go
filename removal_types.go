package iiko

import "fmt"

type RemovalType struct {
	// Identifier. [required]
	ID UUID `json:"id"`

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
	OrganizationIDs []UUID `json:"organizationIds"`
}

type RemovalTypesResponse struct {
	// Operation ID. [required]
	CorrelationID UUID `json:"correlationId"`

	// List of removal types.
	RemovalTypes []RemovalType `json:"paymentTypes"`
}

// Removal types.
//
// iiko API: /api/1/removal_types
func (c *Client) RemovalTypes(req *RemovalTypesRequest, opts ...Option) (*RemovalTypesResponse, error) {
	var (
		onSuccess RemovalTypesResponse
		onError   errorResponse
	)
	if err := c.post(true, "/api/1/removal_types", req, &onSuccess, &onError, opts...); err != nil {
		return nil, err
	}
	if onError.ErrorDescription != "" {
		return nil, fmt.Errorf("iiko error: %q", onError.ErrorDescription)
	}
	return &onSuccess, nil
}
