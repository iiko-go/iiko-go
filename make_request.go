package iiko

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrBadRequest          = errors.New("400 bad request")
	ErrUnauthorized        = errors.New("401 unauthorized")
	ErrRequestTimeout      = errors.New("408 request timeout")
	ErrInternalServerError = errors.New("500 internal server error")
)

func (c *Client) post(requiresAuth bool, endpoint string, body interface{}, onSucces interface{}, onError interface{}, opts ...Option) error {
	if requiresAuth && c.token == "" {
		return ErrUnauthorized
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.baseURL+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	if requiresAuth {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	req.Header.Set("Timeout", defaultTimeout)

	for _, opt := range opts {
		opt.Apply(req)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		if err = json.NewDecoder(resp.Body).Decode(onError); err != nil {
			switch resp.StatusCode {
			case 400:
				return ErrBadRequest
			case 401:
				return ErrUnauthorized
			case 408:
				return ErrRequestTimeout
			case 500:
				return ErrInternalServerError
			}
			return err
		}
		return nil
	}

	if err = json.NewDecoder(resp.Body).Decode(onSucces); err != nil {
		return err
	}
	return nil
}
