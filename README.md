_Library is still under construction. Most of the methods are not implemented yet._
1. [How to use](#how-to-use)
2. [How to contriubte](#how-to-contribute)
3. [Feature matrix](#feature-matrix)
#### How to use
```golang
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/wollzy/iiko-go"
)

func main() {
	client, err := iiko.NewClient("API_LOGIN")
	if err != nil {
		log.Fatalln(err)
	}

	// You can set a custom http.Client for making iikoCloud API request.
	// By default, http.DefaultClient is used.
	client.SetHTTPClient(&http.Client{})

	// You can set a custom timeout for making iikoCloud API request.
	// By default, 15 seconds is used.
	client.SetTimeout(5 * time.Second)

	// You can set a custom refreshTokenInterval for iiko.Client.
	// refreshTokenInterval is the interval at which the iiko.Client wil try to get a new API token.
	// by default, 45 minutes is used.
	client.SetRefreshTokenInterval(30 * time.Minute)

	// make request to iikoCloud API: /api/1/organizations
	res, err := client.Organizations(&iiko.OrganizationsRequest{ReturnAdditionalInfo: true})
	if err != nil {
		// Check if the error is IIKO API Error.
		if iikoError, ok := err.(*iiko.ErrorResponse); ok {
			fmt.Println(iikoError.StatusCode)
			fmt.Println(iikoError.CorrelationID)
			fmt.Println(iikoError.ErrorDescription)
			fmt.Println(iikoError.ErrorField)
			return
		} else {
			log.Fatalln(err)
		}
	}

	// Now we can work with IIKO response.
	for _, organization := range res.Organizations {
		fmt.Println(organization.ID, organization.Country, organization.RestaurantAddress)
	}
}
```
### How to contribute
_In the instructions below you should replace **method_name** or **MethodName** by actual method name that you are writing._
1. Choose a method from [feature matrix](#feature-matrix) that is not implemented yet.
2. Create file **method_name.go** with this template:
```golang
package iiko

type MethodNameRequest struct{}

type MethodNameResponse struct{}

// MethodName description here.
//
// iiko API: /api/1/method_name
func (c *Client) MethodName(req *MethodNameRequest, opts ...Option) (*MethodNameResponse, error) {
	var response MethodNameResponse

	if err := c.post(false, "/api/1/method_name", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
```
3. Put all required types and functions in this one file.
4. Create pull request.
#### Feature matrix
- [x] /access_token
- [x] /organizations
- [x] /cancel_causes
- [x] /combo/get_combos_info
- [x] /commands/status
- [x] /deliveries/order_types
- [x] /discounts
- [x] /nomenclature
- [x] /notifications/send
- [x] /organizations
- [x] /payment_types
- [x] /removal_types
- [x] /stop_lists
- [x] /terminal_groups/is_alive
- [x] /tips_types
- [x] /cities
- [ ] /streets/by_city
- [ ] /deliveries/create
- [ ] /deliveries/update_order_problem
- [ ] /deliveries/update_order_delivery_status
- [ ] /deliveries/update_order_courier
- [ ] /deliveries/add_items
- [ ] /deliveries/close
- [ ] /deliveries/cancel
- [ ] /deliveries/change_complete_before
- [ ] /deliveries/change_delivery_point
- [ ] /deliveries/change_service_type
- [ ] /deliveries/change_payments
- [ ] /deliveries/change_comment
- [ ] /deliveries/print_delivery_bill
- [ ] /deliveries/by_id
- [ ] /deliveries/by_delivery_date_and_status
- [ ] /deliveries/by_revision
- [ ] /deliveries/by_delivery_date_and_phone
- [ ] /deliveries/by_delivery_date_and_source_key_and_filter
- [ ] /deliveries/drafts/by_id
- [ ] /deliveries/drafts/by_filter
- [ ] /deliveries/drafts/save
- [ ] /deliveries/drafts/commit
- [ ] /delivery_restrictions
- [ ] /delivery_restrictions/update
- [ ] /delivery_restrictions/allowed
- [ ] /employees/couriers/locations/by_time_offset
- [ ] /employees/couriers
- [ ] /employees/couriers/by_role
- [ ] /employees/couriers/active_location/by_terminal
- [ ] /employees/couriers/active_location
- [ ] /marketing_sources
- [ ] /order/create
- [ ] /order/by_id
- [ ] /order/by_table
- [ ] /order/add_items
- [ ] /order/close
- [ ] /order/change_payments
- [ ] /reserve/available_organizations
- [ ] /reserve/available_terminal_groups
- [ ] /reserve/available_restaurant_sections
- [ ] /reserve/restaurant_sections_workload
- [ ] /reserve/create
- [ ] /reserve/status_by_id
- [ ] /webhooks/settings
- [ ] /webhooks/update_settings
- [ ] /loyalty/iiko/get_customer
- [ ] /loyalty/iiko/calculate_checkin
- [ ] /loyalty/iiko/get_manual_conditions
- [ ] /combo/calculate_combo_price
- [ ] /regions
- [x] /loyalty/iiko/customer/info
- [x] /loyalty/iiko/customer/card/add
