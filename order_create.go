package iiko

import (
	"time"

	"github.com/google/uuid"
)

type OrderCreateRequest struct {
	// Organization ID.
	OrganizationID uuid.UUID `json:"organizationId"`
	// Front group ID an order must be sent to.
	TerminalGroupID uuid.UUID `json:"terminalGroupId"`
	// Order.
	Order Order `json:"order"`
	// Order creation parameters.
	CreateOrderSettings CreateOrderSettings `json:"createOrderSettings"`
}

type OrderCreateResponse struct {
	// Operation ID.
	CorrelationID uuid.UUID `json:"correlationId"`
	// Order.
	OrderInfo OrderInfo `json:"orderInfo"`
}

type OrderInfo struct {
	// Order ID.
	ID uuid.UUID `json:"id"`
	// POS order ID.
	PosID uuid.UUID `json:"posId"`
	// Order external number.
	ExternalNumber string `json:"externalNumber"`
	// Organization ID.
	// Can be obtained by /api/1/organizations operation.
	OrganizationID uuid.UUID `json:"organizationId"`
	// Timestamp of most recent order change that took place on iikoTransport server.
	Timestamp int `json:"timestamp"`
	// Order creation status. In case of asynchronous creation, it allows to track the instance an order was validated/created in iikoFront.
	CreationStatus OrderCreationStatus `json:"creationStatus"`
	// Order creation error details.
	// Required only if "creationStatus"="Error".
	ErrorInfo ErrorInfo `json:"errorInfo"`
	// Order creation details.
	// Field is filled up if "creationStatus"="Success".
	Order Order `json:"order"`
}

type ErrorInfoCode string

const (
	ErrorInfoCodeCommon                                                                ErrorInfoCode = "Common"
	CommonIllegalDeliveryStatus                                                        ErrorInfoCode = "IllegalDeliveryStatus"
	IllegalDeliveryStatusCustomerNameNotSpecified                                      ErrorInfoCode = "CustomerNameNotSpecified"
	CustomerNameNotSpecifiedProductNotFound                                            ErrorInfoCode = "ProductNotFound"
	ProductNotFoundMarketingSourceNotFound                                             ErrorInfoCode = "MarketingSourceNotFound"
	MarketingSourceNotFoundPaymentTypeNotFound                                         ErrorInfoCode = "PaymentTypeNotFound"
	PaymentTypeNotFoundProductSizeNotFound                                             ErrorInfoCode = "ProductSizeNotFound"
	ProductSizeNotFoundProductGroupNotFound                                            ErrorInfoCode = "ProductGroupNotFound"
	ProductGroupNotFoundOrderNotFound                                                  ErrorInfoCode = "OrderNotFound"
	OrderNotFoundConceptionNotFound                                                    ErrorInfoCode = "ConceptionNotFound"
	ConceptionNotFoundDuplicatedOrderId                                                ErrorInfoCode = "DuplicatedOrderId"
	DuplicatedOrderIdTerminalGroupIdNotDetermined                                      ErrorInfoCode = "TerminalGroupIdNotDetermined"
	TerminalGroupIdNotDeterminedTerminalGroupUnregistered                              ErrorInfoCode = "TerminalGroupUnregistered"
	TerminalGroupUnregisteredInvalidPhone                                              ErrorInfoCode = "InvalidPhone"
	InvalidPhoneModifierDuplicated                                                     ErrorInfoCode = "ModifierDuplicated"
	ModifierDuplicatedProductCanBuyFromCashdesk                                        ErrorInfoCode = "ProductCanBuyFromCashdesk"
	ProductCanBuyFromCashdeskDeliveryOpinionMarkInvalid                                ErrorInfoCode = "DeliveryOpinionMarkInvalid"
	DeliveryOpinionMarkInvalidWrongDeliveryStatusForOpinion                            ErrorInfoCode = "WrongDeliveryStatusForOpinion"
	WrongDeliveryStatusForOpinionOpinionCommentTooLong                                 ErrorInfoCode = "OpinionCommentTooLong"
	OpinionCommentTooLongSurveyItemNotFound                                            ErrorInfoCode = "SurveyItemNotFound"
	SurveyItemNotFoundPaymentTypeCanNotBeUsedAsExternal                                ErrorInfoCode = "PaymentTypeCanNotBeUsedAsExternal"
	PaymentTypeCanNotBeUsedAsExternalAddressNotFound                                   ErrorInfoCode = "AddressNotFound"
	AddressNotFoundHomeNotFound                                                        ErrorInfoCode = "HomeNotFound"
	HomeNotFoundIikonetPaymentAdditionalDataCanNotBeParsed                             ErrorInfoCode = "IikonetPaymentAdditionalDataCanNotBeParsed"
	IikonetPaymentAdditionalDataCanNotBeParsedIikonetPaymentExternalIdNotFound         ErrorInfoCode = "IikonetPaymentExternalIdNotFound"
	IikonetPaymentExternalIdNotFoundIikonetPaymentSumLessThanMarketingDiscount         ErrorInfoCode = "IikonetPaymentSumLessThanMarketingDiscount"
	IikonetPaymentSumLessThanMarketingDiscountDiscountCardNotFound                     ErrorInfoCode = "DiscountCardNotFound"
	DiscountCardNotFoundDiscountCardTypeModeForbidden                                  ErrorInfoCode = "DiscountCardTypeModeForbidden"
	DiscountCardTypeModeForbiddenIikocard5PaymentAdditionalDataCanNotBeParsed          ErrorInfoCode = "Iikocard5PaymentAdditionalDataCanNotBeParsed"
	Iikocard5PaymentAdditionalDataCanNotBeParsedIikocard5PaymentExternalIdNotFound     ErrorInfoCode = "Iikocard5PaymentExternalIdNotFound"
	Iikocard5PaymentExternalIdNotFoundIikocard5PaymentSumLessThanMarketingDiscount     ErrorInfoCode = "Iikocard5PaymentSumLessThanMarketingDiscount"
	Iikocard5PaymentSumLessThanMarketingDiscountIikocard5PaymentCanNotCreateCustomData ErrorInfoCode = "Iikocard5PaymentCanNotCreateCustomData"
	Iikocard5PaymentCanNotCreateCustomDataCourierIdDoesNotExist                        ErrorInfoCode = "CourierIdDoesNotExist"
	CourierIdDoesNotExistCourierDoesNotOwnOrder                                        ErrorInfoCode = "CourierDoesNotOwnOrder"
	CourierDoesNotOwnOrderWrongDeliveryStatus                                          ErrorInfoCode = "WrongDeliveryStatus"
	WrongDeliveryStatusCanNotAssignCourierToOrder                                      ErrorInfoCode = "CanNotAssignCourierToOrder"
	CanNotAssignCourierToOrderUserNotFoundByExternalPassword                           ErrorInfoCode = "UserNotFoundByExternalPassword"
	UserNotFoundByExternalPasswordUserNotFound                                         ErrorInfoCode = "UserNotFound"
	UserNotFoundIikocard51PaymentAdditionalDataCanNotBeParsed                          ErrorInfoCode = "Iikocard51PaymentAdditionalDataCanNotBeParsed"
	Iikocard51PaymentAdditionalDataCanNotBeParsedIikocard51PaymentCredentialNotFound   ErrorInfoCode = "Iikocard51PaymentCredentialNotFound"
	Iikocard51PaymentCredentialNotFoundIikocard51PaymentSearchScopeNotFound            ErrorInfoCode = "Iikocard51PaymentSearchScopeNotFound"
	Iikocard51PaymentSearchScopeNotFoundComboDuplicated                                ErrorInfoCode = "ComboDuplicated"
	ComboDuplicatedInvalidReferenceToCombo                                             ErrorInfoCode = "InvalidReferenceToCombo"
	InvalidReferenceToComboInvalidComboItemsAmount                                     ErrorInfoCode = "InvalidComboItemsAmount"
	InvalidComboItemsAmountInvalidComboItemsGuest                                      ErrorInfoCode = "InvalidComboItemsGuest"
	InvalidComboItemsGuestInvalidReferenceToGuest                                      ErrorInfoCode = "InvalidReferenceToGuest"
	InvalidReferenceToGuestGuestDuplicated                                             ErrorInfoCode = "GuestDuplicated"
	GuestDuplicatedGuestNameNotSpecified                                               ErrorInfoCode = "GuestNameNotSpecified"
	GuestNameNotSpecifiedOrderTypeNotFound                                             ErrorInfoCode = "OrderTypeNotFound"
	OrderTypeNotFoundOrderServiceTypeDoesNotMatchSelfServiceMode                       ErrorInfoCode = "OrderServiceTypeDoesNotMatchSelfServiceMode"
	OrderServiceTypeDoesNotMatchSelfServiceModeDeliveryDateNotSpecified                ErrorInfoCode = "DeliveryDateNotSpecified"
	DeliveryDateNotSpecifiedOrderStatusChangedInIikoFront                              ErrorInfoCode = "OrderStatusChangedInIikoFront"
	OrderStatusChangedInIikoFrontPaymentAdditionalDataTooLong                          ErrorInfoCode = "PaymentAdditionalDataTooLong"
	PaymentAdditionalDataTooLongPaymentSumShouldBePositive                             ErrorInfoCode = "PaymentSumShouldBePositive"
	PaymentSumShouldBePositiveDiscountSumNotSpecified                                  ErrorInfoCode = "DiscountSumNotSpecified"
	DiscountSumNotSpecifiedInvalidDiscountItem                                         ErrorInfoCode = "InvalidDiscountItem"
	InvalidDiscountItemRequestProductPriceIsNotEqualToFrontPrice                       ErrorInfoCode = "RequestProductPriceIsNotEqualToFrontPrice"
	RequestProductPriceIsNotEqualToFrontPriceOrderItemsNotExists                       ErrorInfoCode = "OrderItemsNotExists"
	OrderItemsNotExistsEntityAlreadyInUse                                              ErrorInfoCode = "EntityAlreadyInUse"
	EntityAlreadyInUseDiscountItemPositionNotFound                                     ErrorInfoCode = "DiscountItemPositionNotFound"
	DiscountItemPositionNotFoundDiscountItemDuplicatePositions                         ErrorInfoCode = "DiscountItemDuplicatePositions"
	DiscountItemDuplicatePositionsNonUnqiueOrderItemPosition                           ErrorInfoCode = "NonUnqiueOrderItemPosition"
	NonUnqiueOrderItemPositionEmptyOrderItemPosition                                   ErrorInfoCode = "EmptyOrderItemPosition"
	EmptyOrderItemPositionIncorrectOrderType                                           ErrorInfoCode = "IncorrectOrderType"
	IncorrectOrderTypeIncorrect                                                        ErrorInfoCode = "Incorrect"
	IncorrectTerminalGroupDisabled                                                     ErrorInfoCode = "TerminalGroupDisabled"
	TerminalGroupDisabledOrganizationUnregistered                                      ErrorInfoCode = "OrganizationUnregistered"
	OrganizationUnregisteredOrganizationDisabled                                       ErrorInfoCode = "OrganizationDisabled"
	OrganizationDisabledTooSmallDeliveryDate                                           ErrorInfoCode = "TooSmallDeliveryDate"
	TooSmallDeliveryDateIikoFrontTooOldVersion                                         ErrorInfoCode = "IikoFrontTooOldVersion"
	IikoFrontTooOldVersionInternalServerError                                          ErrorInfoCode = "InternalServerError"
	InternalServerErrorUnknownError                                                    ErrorInfoCode = "UnknownError"
)

type ErrorInfo struct {
	// Error code.
	Code ErrorInfoCode `json:"code"`
	// Nonlocalized message.
	Message string `json:"message"`
	// Localized message.
	Description string `json:"description"`
}

type OrderCreationStatus string

const (
	OrderCreationStatusSuccess    OrderCreationStatus = "Success"
	OrderCreationStatusInProgress OrderCreationStatus = "InProgress"
	OrderCreationStatusError      OrderCreationStatus = "Error"
)

type CreateOrderSettings struct {
	// Timeout in seconds that specifies how much time is given for order to reach iikoFront.
	// After this time, order is nullified if iikoFront doesn't take it. By default - 8 seconds.
	TransportToFrontTimeout int `json:"transportToFrontTimeout"`
}

type Customer struct {
	// Existing customer ID in RMS.
	// If null - the phone number is searched in database, otherwise the new customer is created in RMS.
	ID uuid.UUID `json:"id"`
	// Name of customer.
	// Required for new customers (i.e. if "id" == null) Not required if "id" specified.
	Name string `json:"name"`
	// Last name.
	Surname string `json:"surname"`
	// Comment.
	Comment string `json:"comment"`
	// Date of birth.
	Birthdate time.Time `json:"birthdate"`
	// Email.
	Email string `json:"email"`
	// Whether customer receives order status notification messages.
	ShouldReceiveOrderStatusNotifications bool `json:"shouldReceiveOrderStatusNotifications"`
	// Gender.
	Gender Gender `json:"gender"`
	// Is client in blacklist.
	InBlacklist bool `json:"inBlacklist"`
	// Reason why client was added to blacklist.
	BlacklistReason string `json:"blacklistReason"`
}

type Gender string

const (
	NotSpecified Gender = "NotSpecified"
	Male         Gender = "Male"
	Female       Gender = "Female"
)

type Guests struct {
	// Number of persons.
	Count int `json:"count"`
	// Attribute that shows whether order must be split among guests.
	SplitBetweenPersons bool `json:"splitBetweenPersons"`
}

type Payment struct {
	// Payment type.
	// Can be obtained by /api/1/payment_types operation.
	PaymentTypeID   uuid.UUID       `json:"paymentTypeId"`
	PaymentTypeKind PaymentTypeKind `json:"paymentTypeKind"`
	// Amount due.
	Sum float64 `json:"sum"`
	// Additional payment parameters.
	PaymentAdditionalData PaymentAdditionalData `json:"paymentAdditionalData"`
	// Whether payment item is processed by external payment system (made from outside).
	IsProcessedExternally bool `json:"isProcessedExternally"`
	// Whether the payment item is externally fiscalized.
	IsFiscalizedExternally bool `json:"isFiscalizedExternally"`
}

type SearchScope string

const (
	SearchScopeReserved     SearchScope = "Reserved"
	SearchScopePhone        SearchScope = "Phone"
	SearchScopeCardNumber   SearchScope = "CardNumber"
	SearchScopeCardTrack    SearchScope = "CardTrack"
	SearchScopePaymentToken SearchScope = "PaymentToken"
	SearchScopeFindFaceId   SearchScope = "FindFaceId"
)

type PaymentAdditionalData struct {
	Credential  string          `json:"credential"`
	SearchScope SearchScope     `json:"searchScope"`
	Type        PaymentTypeKind `json:"type"`
}

type Combo struct {
	// Combo ID.
	Id uuid.UUID `json:"id"`
	// Name of combo.
	Name string `json:"name"`
	// Quantity.
	Amount int `json:"amount"`
	// Price of one combo.
	Price float64 `json:"price"`
	// Combo validity ID.
	SourceId uuid.UUID `json:"sourceid"`
	// Combo validity ID.
	ProgramId uuid.UUID `json:"programid"`
}

type OrderItemType string

const (
	OrderItemProduct  OrderItemType = "Product"
	OrderItemCompound OrderItemType = "Compound"
)

type Item struct {
	// Main component.
	PrimaryComponent Component `json:"primaryComponent"`
	// Minor component.
	SecondaryComponent *Component `json:"secondaryComponent"`
	// Indivisible modifiers.
	CommonModifiers []Modifier `json:"commonModifiers"`
	// OrderItemType.
	Type OrderItemType `json:"type"`
	// Quantity.
	Amount int `json:"amount"`
	// Size ID. Required if a stock list item has a size scale.
	ProductSizeID uuid.UUID `json:"productSizeId"`
	// Combo details if combo includes order item.
	ComboInformation ComboInformation `json:"comboInformation"`
	// Comment.
	Comment string `json:"comment"`
}

type ComboInformation struct {
	// Created combo ID.
	ComboID uuid.UUID `json:"comboId"`
	// Action ID that defines combo.
	ComboSourceID uuid.UUID `json:"comboSourceId"`
	// Combo group ID to which item belongs.
	ComboGroupID uuid.UUID `json:"comboGroupId"`
}

type Component struct {
	ProductID  uuid.UUID  `json:"productId"`
	Price      float64    `json:"float64"`
	PositionID uuid.UUID  `json:"positionId"`
	Modifiers  []Modifier `json:"modifiers"`
}

type Modifier struct {
	ProductID uuid.UUID `json:"productId"`
	Amount    int       `json:"amount"`
	// Modifiers group ID (for group modifier). Required for a group modifier.
	// Can be obtained by /api/1/nomenclature operation.
	ProductGroupID uuid.UUID `json:"productGroupId"`
	Price          float64   `json:"float64"`
	PositionID     uuid.UUID `json:"positionId"`
}

type Tip struct {
	// Tips type ID.
	// Can be obtained by /api/1/tips_types operation.
	TipsTypeID uuid.UUID `json:"tipsTypeId"`
	// Payment type.
	// Can be obtained by /api/1/payment_types operation.
	PaymentTypeID   uuid.UUID       `json:"paymentTypeId"`
	PaymentTypeKind PaymentTypeKind `json:"paymentTypeKind"`
	// Amount due.
	Sum float64 `json:"sum"`
	// Additional payment parameters.
	PaymentAdditionalData PaymentAdditionalData `json:"paymentAdditionalData"`
	// Whether payment item is processed by external payment system (made from outside).
	IsProcessedExternally bool `json:"isProcessedExternally"`
	// Whether the payment item is externally fiscalized.
	IsFiscalizedExternally bool `json:"isFiscalizedExternally"`
}

type Waiter struct {
	// ID.
	ID uuid.UUID `json:"id"`
	// Name.
	Name string `json:"name"`
	// Phone.
	Phone string `json:"phone"`
}

type Order struct {
	// Order ID.
	ID uuid.UUID `json:"id"`
	// Order type ID. Can be obtained by /api/1/deliveries/order_types operation.
	OrderTypeID uuid.UUID `json:"orderTypeId"`
	// Order external number.
	ExternalNumber string `json:"externalNumber"`
	// TableIDs. Can be obtained by /api/1/reserve/available_restaurant_sections operation.
	TableIDs []uuid.UUID `json:"tableIds"`
	// Guest.
	Customer Customer `json:"customer"`
	// Guest phone.
	Phone string `json:"phone"`
	// Order status.
	OrderStatus OrderStatus `json:"status"`
	// Order creation date (terminal time zone).
	WhenCreated time.Time `json:"whenCreated"`
	// Order waiter.
	Waiter Waiter `json:"waiter"`
	// Guests information.
	Guests Guests `json:"guests"`
	// Tab name (only for fastfood terminals group in tab mode).
	TabName string `json:"tabName"`
	// Order items.
	Items []Item `json:"items"`
	// Combos included in order.
	Combos []Combo `json:"combos"`
	// Order payment components.
	Payments []Payment `json:"payments"`
	// Order tips components.
	Tips []Tip `json:"tips"`
	// The string key (marker) of the source (partner - api user) that created the order.
	// Needed to limit the visibility of orders for external integration.
	SourceKey string `json:"sourceKey"`
	// Invoice printing time (guest bill time).
	WhenBillPrinted time.Time `json:"whenBillPrinted"`
	// Delivery closing time (Local for delivery terminal).
	WhenClosed time.Time `json:"whenClosed"`
	// Concept.
	Conception Conception `json:"conception"`
	// Discounts/surcharges.
	DiscountsInfo []DiscountInfo `json:"discountsInfo"`
	// Information about iikoCard5.
	IIKOCard5Info IIKOCard5Info `json:"iikoCard5Info"`
}

type Conception struct {
	// ID.
	ID uuid.UUID `json:"id"`
	// Name.
	Name string `json:"name"`
	// Code.
	Code string `json:"code"`
}

type OrderStatus string

const (
	OrderStatusNew     OrderStatus = "New"
	OrderStatusBill    OrderStatus = "Bill"
	OrderStatusClosed  OrderStatus = "Closed"
	OrderStatusDeleted OrderStatus = "Deleted"
)

type IIKOCard5Info struct {
	// Coupon No. that has to be considered when calculating loyalty program.
	Coupon string `json:"coupon"`
	// Information about applied manual conditions.
	ApplicableManualConditions []uuid.UUID `json:"applicableManualConditions"`
}

type Card struct {
	// Track of discount card to be applied to order.
	Track string `json:"track"`
}

type OrderDiscountType string

const (
	OrderDiscountRMS      OrderDiscountType = "RMS"
	OrderDiscountIIKOCard OrderDiscountType = "iikoCard"
)

type OrderDiscount struct {
	// Discount type.
	// Can be obtained by /api/1/discounts operation.
	DiscountTypeID uuid.UUID `json:"discountTypeId"`
	// Discount/surcharge sum.
	Sum float64 `json:"sum"`
	// Order item positions.
	SelectivePositions []uuid.UUID       `json:"selectivePositions"`
	Type               OrderDiscountType `json:"type"`
}

type DiscountInfo struct {
	// Track of discount card to be applied to order.
	Card Card `json:"card"`
	// Discounts/surcharges.
	Discounts []OrderDiscount `json:"discounts"`
}

// OrderCreate ...
//
// iiko API: /api/1/order/create
func (c *Client) OrderCreate(req *OrderCreateResponse, opts ...Option) (*OrderCreateRequest, error) {
	var response OrderCreateRequest

	if err := c.post(true, "/api/1/order/create", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
