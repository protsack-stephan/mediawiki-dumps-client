package main

import (
	"context"
	"fmt"
	"time"

	dumps "github.com/protsack-stephan/mediawiki-dumps-client"
)

func main() {
	client := dumps.NewClient()

	_ = client.PageTitles(context.Background(), "enwikinews", time.Now().UTC(), func(p *dumps.Page) {
		fmt.Println(p)
	})

	date := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)

	_ = client.PageTitlesNs(context.Background(), "enwikinews", date, func(p *dumps.Page) {
		if p.Ns == 6 {
			fmt.Println(p)
		}
	})
}
