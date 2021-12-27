package iiko

import (
	"net/http"
	"strconv"
)

var baseURL string = "https://api-ru.iiko.services"

type Client struct {
	baseURL string
	token   string
	timeout string
	http    *http.Client
}

// SetToken sets IIKO token for Client to use it in Authorization header if required.
func (c *Client) SetToken(token string) {
	c.token = token
}

// SetTimeout sets timeout for Client to use it in Timeout header. By default 15 seconds.
func (c *Client) SetTimeout(seconds int) {
	c.timeout = strconv.FormatInt(int64(seconds), 10)
}

func NewClient() *Client {
	return &Client{
		baseURL: baseURL,
		http:    http.DefaultClient,
	}
}
