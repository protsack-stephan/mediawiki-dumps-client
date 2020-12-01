package dumps

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptions(t *testing.T) {
	opts := newOptions()
	assert.NotNil(t, opts)
	assert.Equal(t, pageTitlesURL, opts.PageTitlesURL)
}
