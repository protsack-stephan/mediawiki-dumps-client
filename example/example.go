package main

import (
	"context"
	"fmt"
	"time"

	dumps "github.com/protsack-stephan/mediawiki-dumps-client"
)

func main() {
	client := dumps.NewCLient()

	titles, err := client.PageTitles(context.Background(), "enwikinews", time.Now().UTC())

	fmt.Println(titles, err)
}
