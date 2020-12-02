package dumps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := NewClient()
	assert.NotNil(t, client)
	assert.NotNil(t, client.httpClient)
	assert.NotNil(t, client.options)
	assert.Equal(t, client.url, dumpsURL)
	assert.Equal(t, client.options.PageTitlesURL, pageTitlesURL)
}
