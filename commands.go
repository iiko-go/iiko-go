package iiko

import (
	"github.com/google/uuid"
)

type CommandsStatusRequest struct {
	OrganizationID uuid.UUID `json:"organizationId"`
	CorrelationID  uuid.UUID `json:"correlationId"`
}

type CommandsStatusType string

const (
	SuccessCommandsStatusType    CommandsStatusType = "Success"
	InProgressCommandsStatusType CommandsStatusType = "InProgress"
	ErrorCommandsStatusType      CommandsStatusType = "Error"
)

type CommandsStatusResponse struct {
	State CommandsStatusType `json:"state"`
}

// Get status of command.
//
// iiko API: /api/1/commands/status
func (c *Client) CommandsStatus(req *CommandsStatusRequest, opts ...Option) (*CommandsStatusResponse, error) {
	var response CommandsStatusResponse

	if err := c.post(true, "/api/1/commands/status", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
