package iiko

import (
	"github.com/google/uuid"
)

type ProductCategoryDiscount struct{}

type Mode string

const (
	Percent     Mode = "Percent"
	FlexibleSum Mode = "FlexibleSum"
	FixedSum    Mode = "FixedSum"
)

type DiscountCardTypeInfo struct {
	// Discount ID in RMS. [required]
	ID uuid.UUID `json:"id"`

	// Discount name. [required]
	Name string `json:"name"`

	// Total discount rate.
	// Ignored if "isCategorisedDiscount" specified. [required]
	Percent float64 `json:"percent"`

	// Whether it is category discount or not.
	// If true, "productCategoryDiscounts" discounts will apply. [required]
	IsCategorisedDiscount bool `json:"isCategorisedDiscount"`

	// Array of objects ProductCategoryDiscount [required]
	ProductCategoryDiscounts []ProductCategoryDiscount `json:"productCategoryDiscounts"`

	// Comment.
	Comment string `json:"comment"`

	// Whether discount allows for selected application to individual items at user's discretion. [required]
	CanBeAppliedSelectively bool `json:"canBeAppliedSelectively"`

	// Minimum order amount required for discount application. If order amount is less than specified threshold, discount does not apply.
	MinOrderSum float64 `json:"minOrderSum"`

	// Enum: "Percent" "FlexibleSum" "FixedSum"`
	// Can be obtained by /api/1/discounts operation. [required]
	Mode Mode `json:"mode"`

	// Fixed amount.
	// Triggers if fixed amount has been specified. [required]
	Sum float64 `json:"sum"`

	// Can be applied by card No.
	// If true, it's enough to enter discount card No. (card swiping not required) [required]
	CanApplyByCardNumber bool `json:"canApplyByCardNumber"`

	// Created manually. [required]
	IsManual bool `json:"isManual"`

	// Executed by card. [required]
	IsCard bool `json:"isCard"`

	// Created automatically. [required]
	IsAutomatic bool `json:"isAutomatic"`

	// IsDeleted.
	IsDeleted bool `json:"isDeleted"`
}

type Discount struct {
	// Organization ID.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationID uuid.UUID `json:"organizationId"`

	// Items for organization. [required]
	Items []DiscountCardTypeInfo `json:"items"`
}

type DiscountsRequest struct {
	// Organization IDs that require discounts return.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationIDs []uuid.UUID `json:"organizationIds"`
}

type DiscountsResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// List of discounts/surcharges. [required]
	Discounts []Discount `json:"discounts"`
}

// Discounts / surcharges.
//
// iiko API: /api/1/discounts
func (c *Client) Discounts(req *DiscountsRequest, opts ...Option) (*DiscountsResponse, error) {
	var response DiscountsResponse

	if err := c.post(true, "/api/1/discounts", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
