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

// PageTitles get list of page titles for project
func (cl *Client) PageTitles(ctx context.Context, dbName string, date time.Time) ([]string, error) {
	url := cl.url + cl.options.PageTitlesURL + "/" + date.Format(dateFormat) + "/" + dbName + "-" + date.Format(dateFormat) + "-all-titles-in-ns-0.gz"
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
