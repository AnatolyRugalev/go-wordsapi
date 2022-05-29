package wordsapi

import (
	"net/http"
	"strings"
)

type options struct {
	baseUrl string
	client  *http.Client
}

const DefaultBaseUrl = "https://wordsapiv1.p.rapidapi.com"

var defaultOptions = options{
	baseUrl: DefaultBaseUrl,
	client:  http.DefaultClient,
}

type Option func(o *options)

func WithBaseUrl(baseUrl string) Option {
	return func(o *options) {
		o.baseUrl = strings.TrimRight(baseUrl, "/")
	}
}

func WithHttpClient(client *http.Client) Option {
	return func(o *options) {
		o.client = client
	}
}
