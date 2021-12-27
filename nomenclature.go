package iiko

import "fmt"

type NomenclatureRequest struct {
	// Organization ID.
	//
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationID UUID `json:"organizationId"`

	// Initial revision. Items list will be received only in case there is a newer revision in the database.
	StartRevision int64 `json:"startRevision"`
}

type Group struct {
	ImageLinks       []string `json:"imageLinks"`
	ParentGroup      string   `json:"parentGroup"`
	Order            int      `json:"order"`
	IsIncludedInMenu bool     `json:"isIncludedInMenu"`
	IsGroupModifier  bool     `json:"isGroupModifier"`
	ID               UUID     `json:"id"`
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	AdditionalInfo   string   `json:"additionalInfo"`
	Tags             []string `json:"tags"`
	IsDeleted        bool     `json:"isDeleted"`
	SeoDescription   string   `json:"seoDescription"`
	SeoText          string   `json:"seoText"`
	SeoKeywords      string   `json:"seoKeywords"`
	SeoTitle         string   `json:"seoTitle"`
}

type ProductCategory struct {
	ID        UUID   `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"isDeleted"`
}

type Price struct {
	CurrentPrice       int    `json:"currentPrice"`
	IsIncludedInMenu   bool   `json:"isIncludedInMenu"`
	NextPrice          int    `json:"nextPrice"`
	NextIncludedInMenu bool   `json:"nextIncludedInMenu"`
	NextDatePrice      string `json:"nextDatePrice"`
}

type SizePrices struct {
	SizeID UUID  `json:"sizeId"`
	Price  Price `json:"price"`
}

type Modifiers struct {
	ID                  string `json:"id"`
	DefaultAmount       int    `json:"defaultAmount"`
	MinAmount           int    `json:"minAmount"`
	MaxAmount           int    `json:"maxAmount"`
	Required            bool   `json:"required"`
	HideIfDefaultAmount bool   `json:"hideIfDefaultAmount"`
	Splittable          bool   `json:"splittable"`
	FreeOfChargeAmount  int    `json:"freeOfChargeAmount"`
}

type ChildModifiers struct {
	ID                  string `json:"id"`
	DefaultAmount       int    `json:"defaultAmount"`
	MinAmount           int    `json:"minAmount"`
	MaxAmount           int    `json:"maxAmount"`
	Required            bool   `json:"required"`
	HideIfDefaultAmount bool   `json:"hideIfDefaultAmount"`
	Splittable          bool   `json:"splittable"`
	FreeOfChargeAmount  int    `json:"freeOfChargeAmount"`
}

type GroupModifiers struct {
	ID                                   string           `json:"id"`
	MinAmount                            int              `json:"minAmount"`
	MaxAmount                            int              `json:"maxAmount"`
	Required                             bool             `json:"required"`
	ChildModifiersHaveMinMaxRestrictions bool             `json:"childModifiersHaveMinMaxRestrictions"`
	ChildModifiers                       []ChildModifiers `json:"childModifiers"`
	HideIfDefaultAmount                  bool             `json:"hideIfDefaultAmount"`
	DefaultAmount                        int              `json:"defaultAmount"`
	Splittable                           bool             `json:"splittable"`
	FreeOfChargeAmount                   int              `json:"freeOfChargeAmount"`
}

type Product struct {
	FatAmount               int              `json:"fatAmount"`
	ProteinsAmount          int              `json:"proteinsAmount"`
	CarbohydratesAmount     int              `json:"carbohydratesAmount"`
	EnergyAmount            int              `json:"energyAmount"`
	FatFullAmount           int              `json:"fatFullAmount"`
	ProteinsFullAmount      int              `json:"proteinsFullAmount"`
	CarbohydratesFullAmount int              `json:"carbohydratesFullAmount"`
	EnergyFullAmount        int              `json:"energyFullAmount"`
	Weight                  int              `json:"weight"`
	GroupID                 string           `json:"groupId"`
	ProductCategoryID       string           `json:"productCategoryId"`
	Type                    string           `json:"type"`
	OrderItemType           string           `json:"orderItemType"`
	ModifierSchemaID        string           `json:"modifierSchemaId"`
	ModifierSchemaName      string           `json:"modifierSchemaName"`
	Splittable              bool             `json:"splittable"`
	MeasureUnit             string           `json:"measureUnit"`
	SizePrices              []SizePrices     `json:"sizePrices"`
	Modifiers               []Modifiers      `json:"modifiers"`
	GroupModifiers          []GroupModifiers `json:"groupModifiers"`
	ImageLinks              []string         `json:"imageLinks"`
	DoNotPrintInCheque      bool             `json:"doNotPrintInCheque"`
	ParentGroup             string           `json:"parentGroup"`
	Order                   int              `json:"order"`
	FullNameEnglish         string           `json:"fullNameEnglish"`
	UseBalanceForSell       bool             `json:"useBalanceForSell"`
	CanSetOpenPrice         bool             `json:"canSetOpenPrice"`
	ID                      string           `json:"id"`
	Code                    string           `json:"code"`
	Name                    string           `json:"name"`
	Description             string           `json:"description"`
	AdditionalInfo          string           `json:"additionalInfo"`
	Tags                    []string         `json:"tags"`
	IsDeleted               bool             `json:"isDeleted"`
	SeoDescription          string           `json:"seoDescription"`
	SeoText                 string           `json:"seoText"`
	SeoKeywords             string           `json:"seoKeywords"`
	SeoTitle                string           `json:"seoTitle"`
}

type Size struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Priority  int    `json:"priority"`
	IsDefault bool   `json:"isDefault"`
}

type NomenclatureResponse struct {
	// Operation ID. [required]
	CorrelationID UUID `json:"correlationId"`

	// Stock list group. [required]
	Groups []Group `json:"groups"`

	// Menu item category. [required]
	ProductCategories []ProductCategory `json:"productCategories"`

	// Menu items and modifiers. [required]
	Products []Product `json:"products"`

	// Item sizes. [required]
	Sizes []Size `json:"sizes"`

	// Items list revision. [required]
	Revision int64 `json:"revision"`
}

// Menu.
//
// iiko API: /api/1/nomenclature
func (c *Client) Nomenclature(req *NomenclatureRequest, opts ...Option) (*NomenclatureResponse, error) {
	var (
		onSuccess NomenclatureResponse
		onError   errorResponse
	)
	if err := c.post(false, "/api/1/nomenclature", req, &onSuccess, &onError, opts...); err != nil {
		return nil, err
	}
	if onError.ErrorDescription != "" {
		return nil, fmt.Errorf("iiko error: %q", onError.ErrorDescription)
	}
	return &onSuccess, nil
}
