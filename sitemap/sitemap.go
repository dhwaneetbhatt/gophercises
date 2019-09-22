package sitemap

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/dhwaneetbhatt/gophercises/link"
)

// Get builds a sitemap of the given website
func Get(website string) {
	pages := get(website)
	for _, page := range pages {
		fmt.Println(page)
	}
}

// get reads a website and returns the list of links on the page
func get(website string) []string {
	resp, err := http.Get(website)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	reqURL := resp.Request.URL
	baseURL := url.URL{
		Scheme: reqURL.Scheme,
		Host:   reqURL.Host,
	}
	base := baseURL.String()
	return filter(hrefs(resp.Body, base), withPrefix(base))
}

// hrefs takes a reader and returns a slice of links
func hrefs(reader io.Reader, base string) []string {
	var ret []string
	links, err := link.Parse(reader)
	if err != nil {
		panic(err)
	}
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}
	return ret
}

// filter filters a slice of links based on the includeLink criteria
func filter(links []string, includeLink func(string) bool) []string {
	var ret []string
	for _, l := range links {
		if includeLink(l) {
			ret = append(ret, l)
		}
	}
	return ret
}

// withPrefix is a filter function that returns true if given string is a prefix of passed string
func withPrefix(prefix string) func(string) bool {
	return func(str string) bool {
		return strings.HasPrefix(str, prefix)
	}
}
