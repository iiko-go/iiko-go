package iiko

import (
	"github.com/google/uuid"
)

type ComboGetCombosInfoRequest struct {
	OrganizationID uuid.UUID `json:"organizationId"`
}

type ComboGetCombosInfoResponse struct {
	ComboSpecifications []ComboSpecification `json:"comboSpecifications"`
	ComboCategories     []ComboCategory      `json:"comboCategories"`
}

type Products struct {
	ProductID               uuid.UUID `json:"productId"`
	SizeID                  uuid.UUID `json:"sizeId"`
	ForbiddenModifiers      []string  `json:"forbiddenModifiers"`
	PriceModificationAmount int       `json:"priceModificationAmount"`
}

type Groups struct {
	ID       uuid.UUID  `json:"id"`
	Name     string     `json:"name"`
	Products []Products `json:"products"`
}

type ComboSpecification struct {
	SourceActionID        uuid.UUID `json:"sourceActionId"`
	CategoryID            uuid.UUID `json:"categoryId"`
	Name                  string    `json:"name"`
	PriceModificationType int       `json:"priceModificationType"`
	PriceModification     int       `json:"priceModification"`
	Groups                []Groups  `json:"groups"`
}

type ComboCategory struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

// Get combos info
//
// iiko API: /api/1/combo/get_combos_info
func (c *Client) ComboGetCombosInfo(req *ComboGetCombosInfoRequest, opts ...Option) (*ComboGetCombosInfoResponse, error) {
	var response ComboGetCombosInfoResponse

	if err := c.post(true, "/api/1/combo/get_combos_info", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
