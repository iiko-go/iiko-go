package iiko

import (
	"net/http"
	"strconv"
	"time"
)

type Option interface {
	Apply(*http.Request)
}

type withCustomTimeout struct {
	timeout time.Duration
}

func (wh withCustomTimeout) Apply(req *http.Request) {
	req.Header.Set("Timeout", strconv.Itoa(int(wh.timeout.Seconds())))
}

// WithCustomTimeout is setting a custom timeout for one API request.
func WithCustomTimeout(t time.Duration) Option {
	return &withCustomTimeout{
		timeout: t,
	}
}
