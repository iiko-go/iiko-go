package iiko

import "net/http"

type Option interface {
	Apply(*http.Request)
}

type withHeader struct {
	key   string
	value string
}

func (wh withHeader) Apply(req *http.Request) {
	req.Header.Add(wh.key, wh.value)
}

func WithHeader(key, value string) Option {
	return &withHeader{
		key:   key,
		value: value,
	}
}
