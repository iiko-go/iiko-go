package iiko

import "errors"

type PaymentType struct{}

type PaymentTypesRequest struct {
	// Organizations IDs which payment types have to be returned.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationIDs []UUID `json:"organizationIds"`
}

type PaymentTypesResponse struct {
	// Operation ID. [required]
	CorrelationID UUID `json:"correlationId"`
	// List of payment types and terminal groups where they are available. [required]
	PaymentTypes []PaymentType `json:"paymentTypes"`
}

// iiko API: /api/1/payment_types
func (c *Client) PaymentTypes(req *PaymentTypesRequest, opts ...Option) (*PaymentTypesResponse, error) {
	var (
		onSuccess PaymentTypesResponse
		onError   errorResponse
	)
	if err := c.post(true, "/api/1/payment_types", req, &onSuccess, &onError, opts...); err != nil {
		return nil, err
	}
	if onError.ErrorDescription != "" {
		return nil, errors.New(onError.ErrorDescription)
	}
	return &onSuccess, nil
}
