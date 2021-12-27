package iiko

import (
	"net/http"
	"strconv"
)

type Option interface {
	Apply(*http.Request)
}

type withCustomTimeout struct {
	seconds int
}

func (wh withCustomTimeout) Apply(req *http.Request) {
	req.Header.Set("Timeout", strconv.Itoa(wh.seconds))
}

func WithCustomTimeout(seconds int) Option {
	return &withCustomTimeout{
		seconds: seconds,
	}
}
