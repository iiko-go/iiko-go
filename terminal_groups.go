package iiko

import (
	"github.com/google/uuid"
)

type TerminalGroupItem struct {
	// Delivery group ID.
	// Can be obtained by /api/1/terminal_groups operation [required]
	ID uuid.UUID `json:"id"`

	// Organization ID.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationID uuid.UUID `json:"organizationId"`

	// Terminal group name. [required]
	Name string `json:"name"`
}

type TerminalGroup struct {
	// Organization ID.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationID uuid.UUID `json:"organizationId"`

	// Items for organization. [required]
	Items []TerminalGroupItem `json:"items"`
}

type TerminalGroupsRequest struct {
	// Organizations IDs for which information is requested.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationIDs []uuid.UUID `json:"organizationIds"`

	// Attribute that shows that response contains disabled terminal groups.
	IncludeDisabled bool `json:"includeDisabled"`
}

type TerminalGroupsResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// List of terminal groups broken down by organizations. [required]
	TerminalGroups []TerminalGroup `json:"terminalGroups"`
}

// Method that returns information on groups of delivery terminals.
//
// iiko API: /api/1/terminal_groups
func (c *Client) TerminalGroups(req *TerminalGroupsRequest, opts ...Option) (*TerminalGroupsResponse, error) {
	var response TerminalGroupsResponse

	if err := c.post(true, "/api/1/terminal_groups", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}

type TerminalGroupsIsAliveRequest struct {
	// Organizations IDs for which information is requested.
	// Can be obtained by /api/1/organizations operation.
	OrganizationIDs []uuid.UUID `json:"organizationIds"`

	// List of terminal groups IDs.
	// Can be obtained by /api/1/terminal_groups operation. [required]
	TerminalGroupIDs []uuid.UUID `json:"terminalGroupIds"`
}

type TerminalGroupAliveInfo struct {
	// Attribute that shows whether a terminal is available to request processing. [required]
	IsAlive bool `json:"isAlive"`

	// ID of front group of terminals.
	// Can be obtained by /api/1/terminal_groups operation. [required]
	TerminalGroupID uuid.UUID `json:"terminalGroupId"`

	// Organizations ID.
	// Can be obtained by /api/1/organizations operation. [required]
	OrganizationID uuid.UUID `json:"organizationId"`
}

type TerminalGroupsIsAliveResponse struct {
	// Operation ID. [required]
	CorrelationID uuid.UUID `json:"correlationId"`

	// Availability attribute of each requested terminal. [required]
	IsAliveStatus []TerminalGroupAliveInfo `json:"isAliveStatus"`
}

// Returns information on availability of group of terminals.
//
// iiko API: /api/1/terminal_groups/is_alive
func (c *Client) TerminalGroupsIsAlive(req *TerminalGroupsIsAliveRequest, opts ...Option) (*TerminalGroupsIsAliveResponse, error) {
	var response TerminalGroupsIsAliveResponse

	if err := c.post(true, "/api/1/terminal_groups/is_alive", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
