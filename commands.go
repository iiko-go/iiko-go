package iiko

import "fmt"

type CommandsStatusRequest struct {
	OrganizationID UUID `json:"organizationId"`
	CorrelationID  UUID `json:"correlationId"`
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
	var (
		onSuccess CommandsStatusResponse
		onError   errorResponse
	)
	if err := c.post(true, "/api/1/commands/status", req, &onSuccess, &onError, opts...); err != nil {
		return nil, err
	}
	if onError.ErrorDescription != "" {
		return nil, fmt.Errorf("iiko error: %q", onError.ErrorDescription)
	}
	return &onSuccess, nil
}
