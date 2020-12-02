package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	dumps "github.com/protsack-stephan/mediawiki-dumps-client"
)

func main() {
	client := dumps.NewBuilder().
		URL("http://new-url.com").
		HTTPClient(&http.Client{}).
		Options(&dumps.Options{}).
		Build()

	// client := dumps.NewCLient()

	titles, err := client.PageTitles(context.Background(), "enwikinews", time.Now().UTC())

	fmt.Println(titles, err)
}
