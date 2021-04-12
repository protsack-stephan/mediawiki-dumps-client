package dumps

import (
	"bufio"
	"compress/gzip"
	"context"
	"fmt"
	"net/http"
	"strconv"
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

func (cl *Client) req(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		return nil, err
	}

	res, err := cl.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code: '%d'", res.StatusCode)
	}

	return res, nil
}

// PageTitles get list of page titles for project in ns 0 (daily)
func (cl *Client) PageTitles(ctx context.Context, dbName string, date time.Time, cb func(p *Page)) error {
	url := fmt.Sprintf("%s%s/%s/%s-%s-all-titles-in-ns-0.gz", cl.url, cl.options.PageTitlesURL, date.Format(dateFormat), dbName, date.Format(dateFormat))
	res, err := cl.req(ctx, url)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	br := bufio.NewReader(res.Body)
	gzr, err := gzip.NewReader(br)

	if err != nil {
		return err
	}

	scn := bufio.NewScanner(gzr)
	scn.Scan()

	for scn.Scan() {
		fields := strings.Fields(scn.Text())

		if len(fields) >= 1 {
			cb(&Page{
				fields[0],
				0,
			})
		}
	}

	return scn.Err()
}

// PageTitelsNs monthly dump of page titles in all namespaces
func (cl *Client) PageTitlesNs(ctx context.Context, dbName string, date time.Time, cb func(*Page)) error {
	url := fmt.Sprintf("%s%s/%s/%s/%s-%s-all-titles.gz", cl.url, cl.options.PageTitlesNsURL, dbName, date.Format(dateFormat), dbName, date.Format(dateFormat))
	res, err := cl.req(ctx, url)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	br := bufio.NewReader(res.Body)
	gzr, err := gzip.NewReader(br)

	if err != nil {
		return err
	}

	scn := bufio.NewScanner(gzr)
	scn.Scan()

	for scn.Scan() {
		fields := strings.Fields(scn.Text())

		if len(fields) >= 2 {
			ns, err := strconv.Atoi(fields[0])

			if err != nil {
				return fmt.Errorf("title: %s, err: %v", fields[1], err)
			}

			cb(&Page{
				fields[1],
				ns,
			})
		}
	}

	return scn.Err()
}
