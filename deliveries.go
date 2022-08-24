package iiko

import (
	"github.com/google/uuid"
)

type OrderServiceType string

const (
	Common            OrderServiceType = "Common"
	DeliveryByCourier OrderServiceType = "DeliveryByCourier"
	DeliveryPickUp    OrderServiceType = "DeliveryPickUp"
)

type OrderTypeItem struct {
	// Order type ID in RMS. [required]
	ID uuid.UUID `json:"id"`

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
	OrganizationID uuid.UUID `json:"organizationId"`

	// Items for organization. [required]
	Items []OrderTypeItem `json:"items"`
}

type DeliveriesOrderTypesRequest struct {
	// Organizations IDs which payment types have to be returned.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationIDs []uuid.UUID `json:"organizationIds"`
}

type DeliveriesOrderTypesResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// List of order types. [required]
	OrderTypes []OrderType `json:"orderTypes"`
}

// Order types.
//
// iiko API: /api/1/deliveries/order_types
func (c *Client) DeliveriesOrderTypes(req *DeliveriesOrderTypesRequest, opts ...Option) (*DeliveriesOrderTypesResponse, error) {
	var response DeliveriesOrderTypesResponse

	if err := c.post(true, "/api/1/deliveries/order_types", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
