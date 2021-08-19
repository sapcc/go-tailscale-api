package client

import (
	"net/http"
	"net/url"
)

const basePath = "/api/v2/"

type Client struct {
	target     url.URL
	apiToken   string
	tailNet    string
	httpClient http.Client
}

func New(url url.URL, apiToken string, tailNet string) (*Client, error) {
	return &Client{
		target:     url,
		apiToken:   apiToken,
		tailNet:    tailNet,
		httpClient: http.Client{},
	}, nil
}
