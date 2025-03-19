package main

import (
	"fmt"
	"sort"
)

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("REPORT for %v\n", baseURL)
	fmt.Println("=============================")

	keys := getSortedKeys(pages)

	for _, key := range keys {
		fmt.Printf("Found %v internal links to %v\n", pages[key], key)
	}
}

func getSortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys
}
