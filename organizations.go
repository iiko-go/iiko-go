package iiko

import (
	"github.com/google/uuid"
)

type ISOCurrency string

const (
	RUB ISOCurrency = "RUB"
	USD ISOCurrency = "USD"
	EUR ISOCurrency = "EUR"
)

type DeliveryServiceType string

const (
	CourierOnly           DeliveryServiceType = "CourierOnly"
	SelfServiceOnly       DeliveryServiceType = "SelfServiceOnly"
	CourierAndSelfService DeliveryServiceType = "CourierAndSelfService"
)

type ResponseType string

const (
	Simple   ResponseType = "Simple"
	Extended ResponseType = "Extended"
)

type Organization struct {
	// Country. [required]
	Country string `json:"country"`

	// Restaurant address. [required]
	RestaurantAddress string `json:"restaurantAddress"`

	// Latitude. [required]
	Latitude float64 `json:"latitude"`

	// Longitude. [required]
	Longitude float64 `json:"longitude"`

	// Regional setting "Use the UAE Addressing System". [required]
	UseUAEAddressingSystem bool `json:"useUaeAddressingSystem"`

	// iikoRms version. [required]
	Version string `json:"version"`

	// ISO currency code (for example: RUB, USD, EUR). [required]
	CurrencyIsoName ISOCurrency `json:"currencyIsoName"`

	// Value rounding of position. [required]
	CurrencyMinimumDenomination float64 `json:"currencyMinimumDenomination"`

	// Country dialing code. [required]
	CountryPhoneCode string `json:"countryPhoneCode"`

	// Require mandatory marketing source input when creating a delivery. [required]
	MarketingSourceRequiredInDelivery bool `json:"marketingSourceRequiredInDelivery"`

	// Default delivery city. [required]
	DefaultDeliveryCityID uuid.UUID `json:"defaultDeliveryCityId"`

	// Delivery cities. [required]
	DeliveryCityIDs []uuid.UUID `json:"deliveryCityIds"`

	// Enum: "CourierOnly" "SelfServiceOnly" "CourierAndSelfService" Delivery type. [required]
	DeliveryServiceType DeliveryServiceType `json:"deliveryServiceType"`

	// Default payment type for CallCenter. [required]
	DefaultCallCenterPaymentTypeID uuid.UUID `json:"defaultCallCenterPaymentTypeId"`

	// Allow text comments for order items (in all restaurant sections). [required]
	OrderItemCommentEnabled bool `json:"orderItemCommentEnabled"`

	// Restaurant`s INN (Taxpayer Identification Number).
	INN string `json:"inn"`

	// [required]
	ResponseType ResponseType `json:"responseType"`

	// Organization ID.
	// Can be obtained by /api/1/organizations operation. [required]
	ID uuid.UUID `json:"id"`

	// Organization name. [required]
	Name string `json:"name"`
}

type OrganizationsRequest struct {
	// Organizations IDs which have to be returned. By default - all organizations from apiLogin.
	// Can be obtained by /api/1/organizations operation.
	OrganizationIDs []uuid.UUID `json:"organizationIds"`

	// A sign whether additional information about the organization should be returned (RMS version, country, restaurantAddress, etc.), or only minimal information should be returned (id and name).
	ReturnAdditionalInfo bool `json:"returnAdditionalInfo"`

	// Attribute that shows that response contains disabled organizations.
	IncludeDisabled bool `json:"includeDisabled"`
}

type OrganizationsResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// List of organizations.
	// Can be obtained by /api/1/organizations operation. [required]
	Organizations []Organization `json:"organizations"`
}

// Returns organizations available to api-login user.
//
// iiko API: /api/1/organizations
func (c *Client) Organizations(req *OrganizationsRequest, opts ...Option) (*OrganizationsResponse, error) {
	var response OrganizationsResponse

	if err := c.post(true, "/api/1/organizations", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
