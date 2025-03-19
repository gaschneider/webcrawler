package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}

	htmlReader := strings.NewReader(htmlBody)
	htmlNode, err := html.Parse(htmlReader)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse HTML: %w", err)
	}

	return checkNodeForURLs(htmlNode, baseURL), nil
}

func checkNodeForURLs(n *html.Node, baseURL *url.URL) []string {
	var urls []string
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				href, err := url.Parse(attr.Val)
				if err != nil {
					fmt.Printf("couldn't parse href '%v': %v\n", attr.Val, err)
					continue
				}
				resolvedURL := baseURL.ResolveReference(href)
				urls = append(urls, resolvedURL.String())
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		urls = append(urls, checkNodeForURLs(c, baseURL)...)
	}

	return urls
}
