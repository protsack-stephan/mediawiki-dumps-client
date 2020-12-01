package dumps

import (
	"net/http"
)

const dumpsURL = "https://dumps.wikimedia.org/"
const dateFormat = "20060102"

// NewCLient create new dumps client
func NewCLient() *Client {
	return &Client{
		dumpsURL,
		new(http.Client),
		newOptions(),
	}
}

// Client for dumps download
type Client struct {
	url        string
	httpClient *http.Client
	options    *Options
}
