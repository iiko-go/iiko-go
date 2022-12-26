package iiko

type CustomerInfoType string

const (
	CustomerInfoTypePhone      CustomerInfoType = "phone"
	CustomerInfoTypeCardTrack  CustomerInfoType = "cardTrack"
	CustomerInfoTypeCardNumber CustomerInfoType = "cardNumber"
	CustomerInfoTypeEmail      CustomerInfoType = "email"
	CustomerInfoTypeId         CustomerInfoType = "id"
)

type CustomerInfoRequest struct {
	Phone          string           `json:"phone,omitempty"`
	CardTrack      string           `json:"cardTrack,omitempty"`
	CardNumber     string           `json:"cardNumber,omitempty"`
	Email          string           `json:"email,omitempty"`
	Id             string           `json:"id,omitempty"`
	Type           CustomerInfoType `json:"type"`
	OrganizationId string           `json:"organizationId"`
}

type SexType int

const (
	SexTypeNotSpecified SexType = 0
	SexTypeMale         SexType = 1
	SexTypeFemale       SexType = 2
)

type ConsentStatus int

const (
	ConsentStatusUnknown ConsentStatus = 0
	ConsentStatusGiven   ConsentStatus = 1
	ConsentStatusRevoked ConsentStatus = 2
)

type CustomerInfoResponse struct {
	Id                            string          `json:"id"`
	ReferrerId                    string          `json:"referrerId"`
	Name                          string          `json:"name"`
	Surname                       string          `json:"surname"`
	MiddleName                    string          `json:"middleName"`
	Comment                       string          `json:"comment"`
	Phone                         string          `json:"phone"`
	CultureName                   string          `json:"cultureName"`
	Birthday                      string          `json:"birthday"`
	Email                         string          `json:"email"`
	Sex                           SexType         `json:"sex"`
	ConsentStatus                 ConsentStatus   `json:"consentStatus"`
	Anonymized                    bool            `json:"anonymized"`
	Cards                         []Card          `json:"cards"`
	Categories                    []Category      `json:"categories"`
	WalletBalances                []WalletBalance `json:"walletBalances"`
	UserData                      string          `json:"userData"`
	ShouldReceivePromoActionsInfo bool            `json:"shouldReceivePromoActionsInfo"`
	ShouldReceiveLoyaltyInfo      bool            `json:"shouldReceiveLoyaltyInfo"`
	ShouldReceiveOrderStatusInfo  bool            `json:"shouldReceiveOrderStatusInfo"`
	PersonalDataConsentFrom       string          `json:"personalDataConsentFrom"`
	PersonalDataConsentTo         string          `json:"personalDataConsentTo"`
	PersonalDataProcessingFrom    string          `json:"personalDataProcessingFrom"`
	PersonalDataProcessingTo      string          `json:"personalDataProcessingTo"`
	IsDeleted                     bool            `json:"isDeleted"`
}

// CustomerInfo ...
//
// iiko API: /api/1/loyalty/iiko/customer/info
func (c *Client) CustomerInfo(req *CustomerInfoRequest, opts ...Option) (*CustomerInfoResponse, error) {
	var response CustomerInfoResponse

	if err := c.post(true, "/api/1/loyalty/iiko/customer/info", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
