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

var pageTitleTestTitles = []string{"🏳️‍🌈", "🏀", "🥊"}
var pageTitlesTestDate = time.Date(2020, 9, 1, 0, 0, 0, 0, time.UTC)

const pageTitlesTestURL = "/pagetitles"
const pageTitlesTestDBName = "test"
const pageTitleTestDate = "20200901"

const pageTitlesTestFile = "test-%s-all-titles-in-ns-0.gz"

func createClientTestServer() http.Handler {
	router := http.NewServeMux()

	path := fmt.Sprintf(pageTitlesTestFile, pageTitleTestDate)
	router.HandleFunc(pageTitlesTestURL+"/"+pageTitleTestDate+"/"+path, func(w http.ResponseWriter, r *http.Request) {
		content, err := ioutil.ReadFile("./stub/" + path)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write(content)
		}
	})

	return router
}

func TestPageTitles(t *testing.T) {
	srv := httptest.NewServer(createClientTestServer())
	defer srv.Close()

	client := NewCLient()
	client.url = srv.URL
	client.options.PageTitlesURL = pageTitlesTestURL

	titles, err := client.PageTitles(context.Background(), pageTitlesTestDBName, pageTitlesTestDate)
	assert.NoError(t, err)
	assert.NotNil(t, titles)

	for _, title := range pageTitleTestTitles {
		assert.Contains(t, titles, title)
	}
}
