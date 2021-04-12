package dumps

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var clientTestTitles = []string{"🏳️‍🌈", "🏀", "🥊"}
var clientTestDate = time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC)

const clientTestDateFormat = "20060102"
const clientTestPageTitleURL = "/pagetitles"
const clientTestDbName = "test"

func createClientTestServer() http.Handler {
	router := http.NewServeMux()

	path := fmt.Sprintf("test-%s-all-titles-in-ns-0.gz", clientTestDate.Format(clientTestDateFormat))
	router.HandleFunc(fmt.Sprintf("%s/%s/%s", clientTestPageTitleURL, clientTestDate.Format(clientTestDateFormat), path), func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("./testdata/" + path)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			_, _ = w.Write(content)
		}
	})

	return router
}

func TestClient(t *testing.T) {
	assert := assert.New(t)

	t.Run("create client", func(t *testing.T) {
		client := NewClient()
		assert.NotNil(client)
		assert.NotNil(client.httpClient)
		assert.NotNil(client.options)
		assert.Equal(client.url, dumpsURL)
		assert.Equal(client.options.PageTitlesURL, pageTitlesURL)
	})

	t.Run("page titles success", func(t *testing.T) {
		srv := httptest.NewServer(createClientTestServer())
		defer srv.Close()

		client := NewClient()
		client.url = srv.URL
		client.options.PageTitlesURL = clientTestPageTitleURL
		titles := map[string]*Page{}

		err := client.PageTitles(context.Background(), clientTestDbName, clientTestDate, func(p *Page) {
			titles[p.Title] = p
		})
		assert.NoError(err)
		assert.NotNil(titles)

		for _, title := range clientTestTitles {
			assert.Contains(titles, title)
		}
	})

	t.Run("page titles error", func(t *testing.T) {
		srv := httptest.NewServer(createClientTestServer())
		defer srv.Close()

		client := NewClient()
		client.url = srv.URL
		client.options.PageTitlesURL = clientTestPageTitleURL
		titles := []*Page{}

		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()

		err := client.PageTitles(ctx, clientTestDbName, clientTestDate, func(p *Page) {
			titles = append(titles, p)
		})
		assert.Contains(err.Error(), context.DeadlineExceeded.Error())
		assert.Equal(0, len(titles))
	})
}
