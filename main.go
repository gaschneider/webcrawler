package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Args == nil || len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(os.Args) > 2 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]

	fmt.Printf("starting crawl of: %v\n", baseURL)

	cfg, err := newConfig(baseURL, 2)
	if err != nil {
		fmt.Printf("Error - newConfig: %v\n", err)
		os.Exit(1)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)

	cfg.wg.Wait()

	for page, count := range cfg.pages {
		fmt.Printf("%v: %v\n", page, count)
	}
}
