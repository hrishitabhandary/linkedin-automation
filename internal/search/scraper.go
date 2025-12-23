package search

import (
	"strings"

	"github.com/go-rod/rod"
)

// CollectProfileURLs extracts LinkedIn profile links from the current page
func CollectProfileURLs(page *rod.Page) []string {
	profileURLs := []string{}
	elements, _ := page.Elements("a.app-aware-link")
	for _, el := range elements {
		href, _ := el.Attribute("href")
		if href != nil && *href != "" && strings.HasPrefix(*href, "/in/") {
			profileURLs = append(profileURLs, "https://www.linkedin.com"+*href)
		}
	}
	return profileURLs
}

// RemoveDuplicates returns unique profile URLs
func RemoveDuplicates(urls []string) []string {
	unique := map[string]bool{}
	result := []string{}
	for _, url := range urls {
		if !unique[url] {
			unique[url] = true
			result = append(result, url)
		}
	}
	return result
}
