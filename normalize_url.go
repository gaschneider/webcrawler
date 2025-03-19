package main

import (
	"fmt"
	neturl "net/url"
	"strings"
)

func normalizeURL(url string) (string, error) {
	parsedUrl, err := neturl.Parse(url)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	path := strings.TrimSuffix(parsedUrl.Path, "/")

	return strings.ToLower(parsedUrl.Host + path), nil
}
