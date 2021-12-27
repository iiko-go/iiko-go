package iiko

import (
	"net/http"
	"strconv"
)

var (
	baseURL string = "https://api-ru.iiko.services"
	// iikoCloud API Timeout in seconds by default
	defaultTimeout = "15"
)

// SetDefaultTimeout sets default Timeout header for all requests. By default 15 seconds.
func SetDefaultTimeout(seconds int) {
	defaultTimeout = strconv.Itoa(seconds)
}

type Client struct {
	baseURL string
	token   string
	timeout string
	http    *http.Client
}

// setToken sets IIKO token for Client to use it in Authorization header if required.
func (c *Client) setToken(token string) {
	c.token = token
}

func NewClient(apiLogin string) (*Client, error) {
	client := &Client{
		baseURL: baseURL,
		http:    http.DefaultClient,
	}

	token, err := client.AccessToken(&AccessTokenRequest{
		ApiLogin: apiLogin,
	})
	if err != nil {
		return nil, err
	}

	client.setToken(token.Token)

	return client, nil
}
