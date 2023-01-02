package iiko

type CardAddRequest struct {
	// uuid
	CustomerId string `json:"customerId"`
	CardTrack  string `json:"cardTrack"`
	CardNumber string `json:"cardNumber"`
	// uuid
	OrganizationId string `json:"organizationId"`
}

type CardAddResponse struct{}

// CardAdd Add new card for customer
//
// iiko API: /api/1/loyalty/iiko/customer/card/add
func (c *Client) CardAdd(req *CardAddRequest, opts ...Option) (*CardAddResponse, error) {
	var response CardAddResponse

	if err := c.post(true, "/api/1/loyalty/iiko/customer/card/add", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
