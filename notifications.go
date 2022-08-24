package iiko

import (
	"fmt"

	"github.com/google/uuid"
)

type NotificationsSendRequest struct {
	// Order source. [required]
	OrderSource string `json:"orderSource"`

	// Order ID. [required]
	OrderID uuid.UUID `json:"orderId"`

	// Additional info about the problem. [required]
	AdditionalInfo string `json:"additionalInfo"`

	// "delivery_attention" [required]
	MessageType string `json:"messageType"`

	// Organization UOC Id. [required]
	OrganizationID uuid.UUID `json:"organizationId"`
}

type NotificationsSendResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`
}

// Send notification to external systems (iikoFront and iikoWeb).
//
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
