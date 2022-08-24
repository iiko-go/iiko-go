package iiko

import (
	"github.com/google/uuid"
)

type TipsType struct {
	// Tips type ID.
	// Can be obtained by /api/1/tips_types operation. [required]
	ID uuid.UUID `json:"id"`

	// Tips type name. [required]
	Name string `json:"name"`

	// Supported organizations IDs. [required]
	// Can be obtained by /api/1/organizations operation.
	OrganizationIDs []uuid.UUID `json:"organizationIds"`

	// Supported order service types. [required]
	OrderServiceTypes []OrderServiceType `json:"orderServiceTypes"`

	// Supported payment types IDs. [required]
	PaymentTypesIDs []uuid.UUID `json:"paymentTypesIds"`
}

type TipsTypesRequest struct{}

type TipsTypesResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// List of tips types for rms group. [required]
	TipsTypes []TipsType `json:"paymentTypes"`
}

// Get tips tipes for api-login`s rms group. Allowed from version 7.7.4.
//
// iiko API: /api/1/tips_types
func (c *Client) TipsTypes(req *TipsTypesRequest, opts ...Option) (*TipsTypesResponse, error) {
	var response TipsTypesResponse

	if err := c.post(true, "/api/1/tips_types", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
