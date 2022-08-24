package iiko

import (
	"fmt"

	"github.com/google/uuid"
)

type StopListsRequest struct {
	OrganizationIds []uuid.UUID `json:"organizationIds"`
}

type StopListsResponse struct {
	CorrelationID          string                  `json:"correlationId"`
	TerminalGroupStopLists []TerminalGroupStopList `json:"terminalGroupStopLists"`
}

type TerminalGroupStopList struct {
	OrganizationID string                      `json:"organizationId"`
	Items          []TerminalGroupStopListItem `json:"items"`
}

type TerminalGroupStopListItem struct {
	Balance   int    `json:"balance"`
	ProductID string `json:"productId"`
}

// Out-of-stock items.
//
// iiko API: /api/1/stop_lists
func (c *Client) StopLists(req *StopListsRequest, opts ...Option) (*StopListsResponse, error) {
	var (
		onSuccess StopListsResponse
		onError   errorResponse
	)
	if err := c.post(false, "/api/1/stop_lists", req, &onSuccess, &onError, opts...); err != nil {
		return nil, err
	}
	if onError.ErrorDescription != "" {
		return nil, fmt.Errorf("iiko error: %q", onError.ErrorDescription)
	}
	return &onSuccess, nil
}
