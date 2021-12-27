package iiko

import (
	"fmt"
)

type NotificationsSendRequest struct {
	// Order source. [required]
	OrderSource string `json:"orderSource"`
	// Order ID. [required]
	OrderID UUID `json:"orderId"`
	// Additional info about the problem. [required]
	AdditionalInfo string `json:"additionalInfo"`
	// "delivery_attention" [required]
	MessageType string `json:"messageType"`
	// Organization UOC Id. [required]
	OrganizationID UUID `json:"organizationId"`
}

type NotificationsSendResponse struct {
	// Operation ID. [required]
	CorrelationID UUID `json:"correlationId"`
}

// iiko API: /api/1/notifications/send
func (c *Client) NotificationsSend(req *NotificationsSendRequest, opts ...Option) (*NotificationsSendResponse, error) {
	var (
		onSuccess NotificationsSendResponse
		onError   errorResponse
	)

	if err := c.post(true, "/api/1/notifications/send", req, &onSuccess, &onError, opts...); err != nil {
		return nil, err
	}

	if onError.ErrorDescription != "" {
		return nil, fmt.Errorf("iiko error: %q", onError.ErrorDescription)
	}
	return &onSuccess, nil
}
