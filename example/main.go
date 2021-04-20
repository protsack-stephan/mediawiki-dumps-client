package main

import (
	"context"
	"fmt"
	"log"
	"time"

	dumps "github.com/protsack-stephan/mediawiki-dumps-client"
)

func main() {
	client := dumps.NewClient()

	_ = client.PageTitles(context.Background(), "enwikinews", time.Now().UTC(), func(p *dumps.Page) {
		fmt.Println(p)
	})

	date := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)

	ns, _ := client.Namespaces(context.Background(), "ukwikinews", date)
	fmt.Println(ns)

	err := client.PageTitlesNs(context.Background(), "ukwikinews", date, func(p *dumps.Page) {
		if p.Ns == 0 {
			fmt.Println(p)
		}
	})

	log.Println(err)
}
