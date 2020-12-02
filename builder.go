package dumps

import "net/http"

// NewBuilder create new instance of builder
func NewBuilder() *ClientBuilder {
	return &ClientBuilder{
		NewClient(),
	}
}

// ClientBuilder is a builder for main client configuration
type ClientBuilder struct {
	client *Client
}

// URL set base url for client
func (cb *ClientBuilder) URL(url string) *ClientBuilder {
	cb.client.url = url
	return cb
}

// HTTPClient pass your own http client
func (cb *ClientBuilder) HTTPClient(httpClient *http.Client) *ClientBuilder {
	cb.client.httpClient = httpClient
	return cb
}

// Options update default options for client
func (cb *ClientBuilder) Options(options *Options) *ClientBuilder {
	cb.client.options = options
	return cb
}

// Build create new client instance
func (cb *ClientBuilder) Build() *Client {
	return cb.client
}
