package link

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represent the link tag in the HTML document
type Link struct {
	Href string
	text string
}

// Parse returns list of parsed links from the HTML document
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	return dfs(doc), nil
}

func dfs(node *html.Node) []Link {
	var links []Link
	if node == nil {
		return links
	}
	if node.Data == "a" {
		link := parseLink(node)
		links = append(links, link)
	}
	links = append(links, dfs(node.FirstChild)...)
	links = append(links, dfs(node.NextSibling)...)
	return links
}

func parseLink(node *html.Node) Link {
	link := Link{}
	attrs := node.Attr
	for _, a := range attrs {
		if a.Key == "href" {
			link.Href = a.Val
		}
	}
	child := node.FirstChild
	for child != nil {
		fmt.Printf("%+v\n", child)
		child = child.NextSibling
	}
	return link
}
