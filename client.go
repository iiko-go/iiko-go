package iiko

import (
	"net/http"
	"time"
)

type Client struct {
	// Channel quit is used to notify that we should stop our JWT-refresh token Ticker.
	quit chan struct{}

	baseURL              string
	apiLogin             string
	token                string
	httpClient           *http.Client
	timeout              time.Duration
	refreshTokenInterval time.Duration
}

// SetTimeout sets default Timeout header for all requests. By default 15 seconds.
func (c *Client) SetTimeout(t time.Duration) {
	c.timeout = t
}

// SetRefreshTokenInterval sets default Timeout header for all requests. By default 15 seconds.
func (c *Client) SetRefreshTokenInterval(t time.Duration) {
	c.refreshTokenInterval = t
}

// SetHTTPClient sets a custom http.Client for making API request to iikoCloud.
func (c *Client) SetHTTPClient(client *http.Client) {
	c.httpClient = client
}

func (c *Client) refreshTokenByInterval() {
	ticker := time.NewTicker(c.refreshTokenInterval)

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
		baseURL:              BaseURL,
		httpClient:           http.DefaultClient,
		apiLogin:             apiLogin,
		timeout:              DefaultTimeout,
		refreshTokenInterval: DefaultRefreshTokenInterval,
	}

	resp, err := client.accessToken(&AccessTokenRequest{
		ApiLogin: client.apiLogin,
	})
	if err != nil {
		return nil, err
	}

	client.token = resp.Token

	go client.refreshTokenByInterval()

	return client, nil
}
