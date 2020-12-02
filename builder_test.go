package dumps

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var builderTestHTTPClient = &http.Client{}
var builderTestOptions = &Options{}

const builderTestURL = "http://builder-test.com"

func TestBuilder(t *testing.T) {
	client := NewBuilder().
		URL(builderTestURL).
		HTTPClient(builderTestHTTPClient).
		Options(builderTestOptions).
		Build()

	assert.NotNil(t, client)
	assert.Equal(t, builderTestURL, client.url)
	assert.Equal(t, builderTestOptions, client.options)
	assert.Equal(t, builderTestHTTPClient, client.httpClient)
}
