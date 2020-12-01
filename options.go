package dumps

const pageTitlesURL = "/other/pagetitles/"

// newOptions create new options struct
func newOptions() *Options {
	return &Options{
		pageTitlesURL,
	}
}

// Options dumps client options
type Options struct {
	PageTitlesURL string
}
