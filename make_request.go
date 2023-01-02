package iiko

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
)

var ErrMissingToken = errors.New("missing API token")

func (c *Client) post(requiresAuth bool, endpoint string, body interface{}, response interface{}, opts ...Option) error {
	if requiresAuth && c.token == "" {
		return ErrMissingToken
	}

	// Marshal json body for request.
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return err
	}

	// Create request.
	req, err := http.NewRequest(http.MethodPost, c.baseURL+endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return err
	}

	// Set headers.
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if requiresAuth {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}
	req.Header.Set("Timeout", strconv.Itoa(int(c.timeout.Seconds())))

	// Apply custom options to request.
	for _, opt := range opts {
		opt.Apply(req)
	}

	// Make request.
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// If error.
	if resp.StatusCode != 200 {
		var errorResponse ErrorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return err
		}
		errorResponse.StatusCode = resp.StatusCode
		return &errorResponse
	}

	// If success.
	if err = json.NewDecoder(resp.Body).Decode(response); err != nil {
		return err
	}

	return nil
}
