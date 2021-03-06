package iiko

import "fmt"

type TipsType struct {
	// Tips type ID.
	// Can be obtained by /api/1/tips_types operation. [required]
	ID UUID `json:"id"`

	// Tips type name. [required]
	Name string `json:"name"`

	// Supported organizations IDs. [required]
	// Can be obtained by /api/1/organizations operation.
	OrganizationIDs []UUID `json:"organizationIds"`

	// Supported order service types. [required]
	OrderServiceTypes []OrderServiceType `json:"orderServiceTypes"`

	// Supported payment types IDs. [required]
	PaymentTypesIDs []UUID `json:"paymentTypesIds"`
}

type TipsTypesRequest struct{}

type TipsTypesResponse struct {
	// Operation ID. [required]
	CorrelationID UUID `json:"correlationId"`

	// List of tips types for rms group. [required]
	TipsTypes []TipsType `json:"paymentTypes"`
}

// Get tips tipes for api-login`s rms group. Allowed from version 7.7.4.
//
// iiko API: /api/1/tips_types
func (c *Client) TipsTypes(req *TipsTypesRequest, opts ...Option) (*TipsTypesResponse, error) {
	var (
		onSuccess TipsTypesResponse
		onError   errorResponse
	)
	if err := c.post(true, "/api/1/tips_types", req, &onSuccess, &onError, opts...); err != nil {
		return nil, err
	}
	if onError.ErrorDescription != "" {
		return nil, fmt.Errorf("iiko error: %q", onError.ErrorDescription)
	}
	return &onSuccess, nil
}
