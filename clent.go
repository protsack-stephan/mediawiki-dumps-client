package dumps

import (
	"bufio"
	"compress/gzip"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const dumpsURL = "https://dumps.wikimedia.org/"
const dateFormat = "20060102"

// NewClient create new dumps client
func NewClient() *Client {
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

// PageTitles get list of page titles for project in ns 0 (daily)
func (cl *Client) PageTitles(ctx context.Context, dbName string, date time.Time) ([]string, error) {
	url := fmt.Sprintf("%s%s/%s/%s-%s-all-titles-in-ns-0.gz", cl.url, cl.options.PageTitlesURL, date.Format(dateFormat), dbName, date.Format(dateFormat))
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return []string{}, err
	}

	res, err := cl.httpClient.Do(req)

	if err != nil {
		return []string{}, err
	}

	if res.StatusCode != http.StatusOK {
		return []string{}, fmt.Errorf("req status '%d'", res.StatusCode)
	}

	defer res.Body.Close()
	br := bufio.NewReader(res.Body)
	gzr, err := gzip.NewReader(br)

	if err != nil {
		return []string{}, err
	}

	body, err := ioutil.ReadAll(gzr)

	if err != nil {
		return []string{}, err
	}

	return strings.Split(strings.TrimSuffix(strings.TrimPrefix(string(body), "page_title\n"), "\n"), "\n"), nil
}
