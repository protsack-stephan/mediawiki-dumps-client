# Mediawiki Dumps Client for GO

This is golang client for accessing Mediawiki API dumps [https://dumps.wikimedia.org/](https://dumps.wikimedia.org/). Right now we are only supporting [https://dumps.wikimedia.org/other/pagetitles/](https://dumps.wikimedia.org/other/pagetitles/) but that can change in the future.

Small example of titles for English Wikinews:
```go
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
```

If you need to change the default configuration you can use client builder:
```go
client := dumps.NewBuilder().
		URL("http://new-url.com").
		HTTPClient(&http.Client{}).
		Options(&dumps.Options{}).
		Build()

date := time.Date(time.Now().Year(), time.Now().Month(), 1, 0, 0, 0, 0, time.UTC)

_ = client.PageTitlesNs(context.Background(), "enwikinews", date, func(p *dumps.Page) {
	if p.Ns == 6 {
		fmt.Println(p)
	}
})
```

### Note that for big projects you might need lots of RAM to keep all the data in memory.