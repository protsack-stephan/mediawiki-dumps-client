package main

import (
	"context"
	"log"
	"time"

	dumps "github.com/protsack-stephan/mediawiki-dumps-client"
)

func main() {
	client := dumps.NewClient()

	err := client.PageTitles(context.Background(), "enwikinews", time.Now().UTC(), func(p *dumps.Page) {
		log.Println(p)
	})

	if err != nil {
		log.Println(err)
	}

	date := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)

	ns, err := client.Namespaces(context.Background(), "ukwikinews", date)

	if err != nil {
		log.Println(err)
	} else {
		log.Println(ns)
	}

	err = client.PageTitlesNs(context.Background(), "ukwikinews", date, func(p *dumps.Page) {
		if p.Ns == 0 {
			log.Println(p)
		}
	})

	if err != nil {
		log.Println(err)
	}
}
