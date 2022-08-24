package iiko

import (
	"github.com/google/uuid"
)

type StopListsRequest struct {
	OrganizationIDs []uuid.UUID `json:"organizationIds"`
}

type StopListsResponse struct {
	CorrelationID          uuid.UUID               `json:"correlationId"`
	TerminalGroupStopLists []TerminalGroupStopList `json:"terminalGroupStopLists"`
}

type TerminalGroupStopList struct {
	OrganizationID uuid.UUID                   `json:"organizationId"`
	Items          []TerminalGroupStopListItem `json:"items"`
}

type TerminalGroupStopListItem struct {
	Balance   int       `json:"balance"`
	ProductID uuid.UUID `json:"productId"`
}

// Out-of-stock items.
//
// iiko API: /api/1/stop_lists
func (c *Client) StopLists(req *StopListsRequest, opts ...Option) (*StopListsResponse, error) {
	var response StopListsResponse

	if err := c.post(false, "/api/1/stop_lists", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
