package iiko

import "github.com/google/uuid"

type CitiesRequest struct {
	OrganizationIDs []uuid.UUID `json:"organizationIds"`
}

type CitiesResponse struct {
	CorrelationID uuid.UUID  `json:"correlationId"`
	Cities        []CityItem `json:"cities"`
}

type CityItem struct {
	OrganizationID uuid.UUID `json:"organizationId"`
	Items          []City    `json:"items"`
}

type City struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	ExternalRevision int       `json:"externalRevision"`
	IsDeleted        bool      `json:"isDeleted"`
	ClassifierID     uuid.UUID `json:"classifierId"`
	AdditionalInfo   string    `json:"additionalInfo"`
}

// Cities ...
//
// iiko API: /api/1/cities
func (c *Client) Cities(req *CitiesRequest, opts ...Option) (*CitiesResponse, error) {
	var response CitiesResponse

	if err := c.post(true, "/api/1/cities", req, &response, opts...); err != nil {
		return nil, err
	}

	return &response, nil
}
