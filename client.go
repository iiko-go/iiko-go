package iiko

import (
	"net/http"
	"time"
)

type Client struct {
	// Used to notify that we should stop our JWT-refresh token Ticker.
	quit     chan struct{}
	baseURL  string
	apiLogin string
	token    string
	http     *http.Client
}

func NewClient(apiLogin string) (*Client, error) {
	client := &Client{
		baseURL:  baseURL,
		http:     http.DefaultClient,
		apiLogin: apiLogin,
	}

	resp, err := client.accessToken(&AccessTokenRequest{
		ApiLogin: client.apiLogin,
	})
	if err != nil {
		return nil, err
	}

	client.token = resp.Token

	go client.refreshTokenCronJob()

	return client, nil
}

func (c *Client) Close() {
	close(c.quit)
}

func (c *Client) refreshTokenCronJob() {
	ticker := time.NewTicker(45 * time.Minute)

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
