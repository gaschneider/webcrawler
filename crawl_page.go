package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	parsedCurrUrl, err := url.Parse(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	if parsedCurrUrl.Host != cfg.baseURL.Host {
		fmt.Printf("Error - crawlPage: URL '%s' is not on the same domain as '%s'\n", rawCurrentURL, cfg.baseURL.Host)
		return
	}

	normalizedUrl, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't normalize URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	isFirst := cfg.addPageVisit(normalizedUrl)
	if !isFirst {
		return
	}

	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't get HTML for URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	fmt.Printf("Crawling %v\n", rawCurrentURL)

	links, err := getURLsFromHTML(htmlBody, cfg.baseURL.String())
	if err != nil {
		fmt.Printf("Error - crawlPage: couldn't get URLs from HTML for URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	for _, link := range links {
		cfg.wg.Add(1)
		go cfg.crawlPage(link)
	}
}
