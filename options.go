package dumps

const pageTitlesURL = "/other/pagetitles/"
const pageTitlesNsURL = ""
const namespacesURL = ""

// newOptions create new options struct
func newOptions() *Options {
	return &Options{
		pageTitlesURL,
		pageTitlesNsURL,
		namespacesURL,
	}
}

// Options dumps client options
type Options struct {
	PageTitlesURL   string
	PageTitlesNsURL string
	NamespacesURL   string
}
