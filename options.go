package dumps

const pageTitlesURL = "/other/pagetitles/"
const pageTitlesNsURL = ""

// newOptions create new options struct
func newOptions() *Options {
	return &Options{
		pageTitlesURL,
		pageTitlesNsURL,
	}
}

// Options dumps client options
type Options struct {
	PageTitlesURL   string
	PageTitlesNsURL string
}
