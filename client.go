package iiko

import (
	"net/http"
	"strconv"
	"time"
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
	quit     chan struct{}
	baseURL  string
	apiLogin string
	token    string
	timeout  string
	http     *http.Client
}

// setToken sets IIKO token for Client to use it in Authorization header if required.
func (c *Client) setToken(token string) {
	c.token = token
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

	client.setToken(resp.Token)

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
			_ = c.refreshToken()
		case <-c.quit:
			ticker.Stop()
			return
		}
	}
}

func (c *Client) refreshToken() error {
	resp, err := c.accessToken(&AccessTokenRequest{
		ApiLogin: c.apiLogin,
	})
	if err != nil {
		return err
	}

	c.setToken(resp.Token)

	return nil
}
