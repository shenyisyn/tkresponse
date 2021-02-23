package jtplugin

import (
	"context"
	"net/http"
)

type Config struct {
}

func CreateConfig() *Config {
	return &Config{}
}

type JtResponse struct {
	next http.Handler
	name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &JtResponse{
		next: next,
		name: name,
	}, nil
}

func (e *JtResponse) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Add("Author", "jtthink.com")
	e.next.ServeHTTP(rw, req)
}
