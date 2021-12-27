package iiko

type errorResponse struct {
	// Operation ID [required]
	CorrelationID UUID `json:"correlationId"`
	// Error text [required]
	ErrorDescription string `json:"errorDescription"`
	// Error code.
	ErrorField string `json:"error"`
}

// iiko API: /api/1/nomenclature
func Nomenclature() {}

// iiko API: /api/1/stop_lists
func StopLists() {}

// iiko API: /api/1/combo/get_combos_info
func ComboGetCombosInfo() {}

// iiko API: /api/1/combo/calculate_combo_price
func ComboCalculateComboPrice() {}

// iiko API: /api/1/commands/status
func CommandsStatus() {}

// iiko API: /api/1/regions
func Regions() {}

// iiko API: /api/1/cities
func Cities() {}

// iiko API: /api/1/streets/by_city
func StreetsByCity() {}

// iiko API: /api/1/deliveries/create
func DeliveriesCreate() {}

// iiko API: /api/1/deliveries/update_order_problem
func DeliveriesUpdateOrderProblem() {}

// iiko API: /api/1/deliveries/update_order_delivery_status
func DeliveriesUpdateOrderDeliveryStatus() {}

// iiko API: /api/1/deliveries/update_order_courier
func DeliveriesUpdateOrderCourier() {}

// iiko API: /api/1/deliveries/add_items
func DeliveriesAddItems() {}

// iiko API: /api/1/deliveries/close
func DeliveriesClose() {}

// iiko API: /api/1/deliveries/cancel
func DeliveriesCancel() {}

// iiko API: /api/1/deliveries/change_complete_before
func DeliveriesChangeCompleteBefore() {}

// iiko API: /api/1/deliveries/change_delivery_point
func DeliveriesChangeDeliveryPoint() {}

// iiko API: /api/1/deliveries/change_service_type
func DeliveriesChangeServiceType() {}

// iiko API: /api/1/deliveries/change_payments
func DeliveriesChangePayments() {}

// iiko API: /api/1/deliveries/change_comment
func DeliveriesChangeComment() {}

// iiko API: /api/1/deliveries/print_delivery_bill
func DeliveriesPrintDeliveryBill() {}

// iiko API: /api/1/deliveries/by_id
func DeliveriesByID() {}

// iiko API: /api/1/deliveries/by_delivery_date_and_status
func DeliveriesByDeliveryDateAndStatus() {}

// iiko API: /api/1/deliveries/by_revision
func DeliveriesByRevision() {}

// iiko API: /api/1/deliveries/by_delivery_date_and_phone
func DeliveriesByDeliveryDateAndPhone() {}

// iiko API: /api/1/deliveries/by_delivery_date_and_source_key_and_filter
func DeliveriesByDeliveryDateAndSourceKeyAndFilter() {}

// iiko API: /api/1/deliveries/drafts/by_id
func DeliveriesDraftsByID() {}

// iiko API: /api/1/deliveries/drafts/by_filter
func DeliveriesByFilter() {}

// iiko API: /api/1/deliveries/drafts/save
func DeliveriesSave() {}

// iiko API: /api/1/deliveries/drafts/commit
func DeliveriesCommit() {}

// iiko API: /api/1/delivery_restrictions
func DeliveryRestrictions() {}

// iiko API: /api/1/delivery_restrictions/update
func Update() {}

// iiko API: /api/1/delivery_restrictions/allowed
func Allowed() {}

// iiko API: /api/1/employees/couriers/locations/by_time_offset
func ByTimeOffset() {}

// iiko API: /api/1/employees/couriers
func Couriers() {}

// iiko API: /api/1/employees/couriers/by_role
func ByRole() {}

// iiko API: /api/1/employees/couriers/active_location/by_terminal
func ByTerminal() {}

// iiko API: /api/1/employees/couriers/active_location
func ActiveLocation() {}

// iiko API: /api/1/marketing_sources
func MarketingSources() {}

// iiko API: /api/1/order/create
func OrderCreate() {}

// iiko API: /api/1/order/by_id
func OrderByID() {}

// iiko API: /api/1/order/by_table
func OrderByTable() {}

// iiko API: /api/1/order/add_items
func OrderAddItems() {}

// iiko API: /api/1/order/close
func OrderClose() {}

// iiko API: /api/1/order/change_payments
func OrderChangePayments() {}

// iiko API: /api/1/reserve/available_organizations
func ReserveAvailableOrganizations() {}

// iiko API: /api/1/reserve/available_terminal_groups
func ReserveAvailableTerminalGroups() {}

// iiko API: /api/1/reserve/available_restaurant_sections
func ReserveAvailableRestaurantSections() {}

// iiko API: /api/1/reserve/restaurant_sections_workload
func ReserveRestaurantSectionsWorkload() {}

// iiko API: /api/1/reserve/create
func ReserveCreate() {}

// iiko API: /api/1/reserve/status_by_id
func ReserveStatusByID() {}

// iiko API: /api/1/webhooks/settings
func WebhooksSettings() {}

// iiko API: /api/1/webhooks/update_settings
func WebhooksUpdateSettings() {}

// iiko API: /api/1/loyalty/iiko/get_customer
func LoyaltyGetCustomer() {}

// iiko API: /api/1/loyalty/iiko/calculate_checkin
func LoyaltyCalculateCheckin() {}

// iiko API: /api/1/loyalty/iiko/get_manual_conditions
func LoyaltyGetManualConditions() {}
