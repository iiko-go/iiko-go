package iiko

import (
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
	var response NotificationsSendResponse

	if err := c.post(true, "/api/1/notifications/send", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
