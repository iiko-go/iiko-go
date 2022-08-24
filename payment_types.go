package iiko

import (
	"github.com/google/uuid"
)

type PaymentProcessingType string

const (
	ExternalProcessingType PaymentProcessingType = "External"
	InternalProcessingType PaymentProcessingType = "Internal"
	BothProcessingType     PaymentProcessingType = "Both"
)

type PaymentTypeKind string

const (
	UnknownPaymentTypeKind  PaymentTypeKind = "Unknown"
	CashPaymentTypeKind     PaymentTypeKind = "Cash"
	CardPaymentTypeKind     PaymentTypeKind = "Card"
	CreditPaymentTypeKind   PaymentTypeKind = "Credit"
	WriteoffPaymentTypeKind PaymentTypeKind = "Writeoff"
	VoucherPaymentTypeKind  PaymentTypeKind = "Voucher"
	ExternalPaymentTypeKind PaymentTypeKind = "External"
	IikoCardPaymentTypeKind PaymentTypeKind = "IikoCard"
)

type PaymentType struct {
	// Payment type ID
	ID uuid.UUID `json:"id"`

	// Payment type code
	Code string `json:"code"`

	// Payment type name
	Name string `json:"name"`

	// Payment type comment
	Comment string `json:"comment"`

	// Combinability attribute
	Combinable bool `json:"combinable"`

	// External system revision number.
	ExternalRevision int64 `json:"externalRevision"`

	// Array of marketing campaigns associated with iikoCard5 payment type applicable to this organization.
	ApplicableMarketingCampaigns []uuid.UUID `json:"applicableMarketingCampaigns"`

	// IsDeleted attribute of payment type.
	IsDeleted bool `json:"isDeleted"`

	// If true, payment type is fiscal and bill will be printed.
	PrintCheque bool `json:"printCheque"`

	// Describes operation processing type. "External" "Internal" "Both"
	PaymentProcessingType PaymentProcessingType `json:"paymentProcessingType"`

	// Payment type category.
	PaymentTypeKind PaymentTypeKind `json:"paymentTypeKind"`

	// Terminal groups where this payment type is available.
	TerminalGroups []TerminalGroup `json:"terminalGroups"`
}

type PaymentTypesRequest struct {
	// Organizations IDs which payment types have to be returned.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationIDs []uuid.UUID `json:"organizationIds"`
}

type PaymentTypesResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// List of payment types and terminal groups where they are available. [required]
	PaymentTypes []PaymentType `json:"paymentTypes"`
}

// Payment types.
//
// iiko API: /api/1/payment_types
func (c *Client) PaymentTypes(req *PaymentTypesRequest, opts ...Option) (*PaymentTypesResponse, error) {
	var response PaymentTypesResponse

	if err := c.post(true, "/api/1/payment_types", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
