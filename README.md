# Mediawiki Dumps Client for GO

This is golang client for accessing Mediawiki API dumps [https://dumps.wikimedia.org/](https://dumps.wikimedia.org/). Right now we are only supporting [https://dumps.wikimedia.org/other/pagetitles/](https://dumps.wikimedia.org/other/pagetitles/) but that can change in the future.

Small example of titles for English Wikinews:
```go
client := dumps.NewCLient()

titles, err := client.PageTitles(context.Background(), "enwikinews", time.Now().UTC())

fmt.Println(titles, err)
```

If you need to change the default configuration you can use client builder:
```go
client := dumps.NewBuilder().
		URL("http://new-url.com").
		HTTPClient(&http.Client{}).
		Options(&dumps.Options{}).
		Build()

titles, err := client.PageTitles(context.Background(), "enwikinews", time.Now().UTC())

fmt.Println(titles, err)
```

### Note that for big projects you might need lots of ram to keep the data in memory.