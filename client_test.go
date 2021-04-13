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
const clientTestPageTitlesNsURL = "/alltitles"
const clientTestDbName = "test"

func createClientTestServer() http.Handler {
	router := http.NewServeMux()

	titles := fmt.Sprintf("test-%s-all-titles-in-ns-0.gz", clientTestDate.Format(clientTestDateFormat))
	router.HandleFunc(fmt.Sprintf("%s/%s/%s", clientTestPageTitleURL, clientTestDate.Format(clientTestDateFormat), titles), func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile(fmt.Sprintf("./testdata/%s", titles))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			_, _ = w.Write(content)
		}
	})

	allTitles := fmt.Sprintf("test-%s-all-titles.gz", clientTestDate.Format(clientTestDateFormat))
	router.HandleFunc(fmt.Sprintf("%s/%s/%s/%s", clientTestPageTitlesNsURL, clientTestDbName, clientTestDate.Format(clientTestDateFormat), allTitles), func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile(fmt.Sprintf("./testdata/%s", allTitles))

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
	srv := httptest.NewServer(createClientTestServer())
	defer srv.Close()

	t.Run("create client", func(t *testing.T) {
		client := NewClient()
		assert.NotNil(client)
		assert.NotNil(client.httpClient)
		assert.NotNil(client.options)
		assert.Equal(client.url, dumpsURL)
		assert.Equal(client.options.PageTitlesURL, pageTitlesURL)
	})

	t.Run("page titles success", func(t *testing.T) {
		client := NewClient()
		client.url = srv.URL
		client.options.PageTitlesURL = clientTestPageTitleURL
		titles := map[string]*Page{}

		err := client.PageTitles(context.Background(), clientTestDbName, clientTestDate, func(p *Page) {
			titles[p.Title] = p
		})
		assert.NoError(err)

		for _, title := range clientTestTitles {
			assert.Contains(titles, title)
		}
	})

	t.Run("page titles error", func(t *testing.T) {
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

	t.Run("page titles ns success", func(t *testing.T) {
		client := NewClient()
		client.url = srv.URL
		client.options.PageTitlesNsURL = clientTestPageTitlesNsURL
		titles := map[string]*Page{}

		err := client.PageTitlesNs(context.Background(), clientTestDbName, clientTestDate, func(p *Page) {
			titles[p.Title] = p
		})
		assert.NoError(err)

		for _, title := range clientTestTitles {
			assert.Contains(titles, title)
		}
	})

	t.Run("page titles ns error", func(t *testing.T) {
		client := NewClient()
		client.url = srv.URL
		client.options.PageTitlesNsURL = clientTestPageTitlesNsURL
		titles := []*Page{}

		ctx, cancel := context.WithTimeout(context.Background(), time.Microsecond*1)
		defer cancel()

		err := client.PageTitlesNs(ctx, clientTestDbName, clientTestDate, func(p *Page) {
			titles = append(titles, p)
		})
		assert.Contains(err.Error(), context.DeadlineExceeded.Error())
		assert.Equal(0, len(titles))
	})
}
