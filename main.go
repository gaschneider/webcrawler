package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if os.Args == nil || len(os.Args) < 2 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(os.Args) > 4 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := os.Args[1]
	maxConcurrency := 2
	maxPages := 10

	if len(os.Args) == 3 {
		if mc, err := strconv.Atoi(os.Args[2]); err == nil {
			maxConcurrency = mc
		} else {
			fmt.Printf("Invalid maxConcurrency value: %v\n", os.Args[2])
			os.Exit(1)
		}
	}

	if len(os.Args) == 4 {
		if mp, err := strconv.Atoi(os.Args[3]); err == nil {
			maxPages = mp
		} else {
			fmt.Printf("Invalid maxConcurrency value: %v\n", os.Args[3])
			os.Exit(1)
		}
	}

	fmt.Printf("starting crawl of: %v\n", baseURL)

	cfg, err := newConfig(baseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("Error - newConfig: %v\n", err)
		os.Exit(1)
	}

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)

	cfg.wg.Wait()

	printReport(cfg.pages, baseURL)
}
