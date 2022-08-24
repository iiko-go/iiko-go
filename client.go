package iiko

import (
	"net/http"
	"time"
)

type Client struct {
	// Channel quit is used to notify that we should stop our JWT-refresh token Ticker.
	quit chan struct{}

	baseURL             string
	apiLogin            string
	token               string
	httpClient          *http.Client
	timeout             time.Duration
	refreshTokenTimeout time.Duration
}

// SetTimeout sets default Timeout header for all requests. By default 15 seconds.
func (c *Client) SetTimeout(t time.Duration) {
	c.timeout = t
}

// SetRefreshTokenTimeout sets default Timeout header for all requests. By default 15 seconds.
func (c *Client) SetRefreshTokenTimeout(t time.Duration) {
	c.refreshTokenTimeout = t
}

// SetHTTPClient sets a custom http.Client for making API request to iikoCloud.
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

//
func (c *Client) refreshTokenByTimeout() {
	ticker := time.NewTicker(c.refreshTokenTimeout)

	c.quit = make(chan struct{})

	for {
		select {
		case <-ticker.C:
			// get new JWT token from iikoCloud API ang ignore error
			resp, _ := c.accessToken(&AccessTokenRequest{
				ApiLogin: c.apiLogin,
			})

			c.token = resp.Token

		case <-c.quit:
			ticker.Stop()
			return
		}
	}
}

func (c *Client) Close() {
	close(c.quit)
}

func NewClient(apiLogin string) (*Client, error) {
	client := &Client{
		baseURL:             BaseURL,
		httpClient:          http.DefaultClient,
		apiLogin:            apiLogin,
		timeout:             DefaultTimeout,
		refreshTokenTimeout: DefaultRefreshTokenTimeout,
	}

	resp, err := client.accessToken(&AccessTokenRequest{
		ApiLogin: client.apiLogin,
	})
	if err != nil {
		return nil, err
	}

	client.token = resp.Token

	go client.refreshTokenByTimeout()

	return client, nil
}
