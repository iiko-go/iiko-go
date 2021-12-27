package iiko

import "errors"

type OrderServiceType string

const (
	Common            OrderServiceType = "Common"
	DeliveryByCourier OrderServiceType = "DeliveryByCourier"
	DeliveryPickUp    OrderServiceType = "DeliveryPickUp"
)

type OrderTypeItem struct {
	// Order type ID in RMS. [required]
	ID UUID `json:"id"`
	// Order type name. [required]
	Name string `json:"name"`
	// Enum: Common, DeliveryByCourier, DeliveryPickUp Service type. [required]
	OrderServiceType OrderServiceType `json:"orderServiceType"`
	// IsDeleted attribute of order type.
	IsDeleted bool `json:"isDeleted"`
	// External system revision number.
	ExternalRevision int64 `json:"externalRevision"`
}

type OrderType struct {
	// Organization ID.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationID UUID `json:"organizationId"`
	// Items for organization. [required]
	Items []OrderTypeItem `json:"items"`
}

type DeliveriesOrderTypesRequest struct {
	// Organizations IDs which payment types have to be returned.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationIDs []UUID `json:"organizationIds"`
}

type DeliveriesOrderTypesResponse struct {
	// Operation ID. [required]
	CorrelationID UUID `json:"correlationId"`
	// List of order types. [required]
	OrderTypes []OrderType `json:"orderTypes"`
}

// iiko API: /api/1/deliveries/order_types
func (c *Client) DeliveriesOrderTypes(req *DeliveriesOrderTypesRequest, opts ...Option) (*DeliveriesOrderTypesResponse, error) {
	var (
		onSuccess DeliveriesOrderTypesResponse
		onError   errorResponse
	)
	if err := c.post(true, "/api/1/deliveries/order_types", req, &onSuccess, &onError, opts...); err != nil {
		return nil, err
	}
	if onError.ErrorDescription != "" {
		return nil, errors.New(onError.ErrorDescription)
	}
	return &onSuccess, nil
}
