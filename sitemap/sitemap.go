package sitemap

import (
	"encoding/xml"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/dhwaneetbhatt/gophercises/link"
)

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

// Get builds a sitemap of the given website
func Get(website string, maxDepth int, writer io.Writer) error {
	pages := bfs(website, maxDepth)
	return writeToXML(pages, writer)
}

// writeToXML writes the pages to Sitemap XML format
func writeToXML(pages []string, writer io.Writer) error {
	toXML := urlset{Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	for _, page := range pages {
		toXML.Urls = append(toXML.Urls, loc{page})
	}
	writer.Write([]byte(xml.Header))
	encoder := xml.NewEncoder(writer)
	encoder.Indent("", "  ")
	if err := encoder.Encode(toXML); err != nil {
		return err
	}
	return nil
}

// bfs does a BFS of the website
func bfs(website string, maxDepth int) []string {
	visited := make(map[string]bool)
	var queue map[string]bool
	nextQueue := map[string]bool{
		website: true,
	}
	for i := 0; i <= maxDepth; i++ {
		queue, nextQueue = nextQueue, make(map[string]bool)
		if len(queue) == 0 {
			break
		}
		for url := range queue {
			if visited[url] {
				continue
			}
			visited[url] = true
			for _, link := range get(url) {
				if !visited[link] {
					nextQueue[link] = true
				}
			}
		}
	}
	ret := make([]string, 0, len(visited))
	for url := range visited {
		ret = append(ret, url)
	}
	return ret
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
