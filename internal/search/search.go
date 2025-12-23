package search

import (
	"net/url"
)

// BuildSearchURL creates a LinkedIn people search URL
func BuildSearchURL(title, location, company, keywords string) string {
	base := "https://www.linkedin.com/search/results/people/?"

	params := url.Values{}
	if title != "" {
		params.Add("keywords", title)
	}
	if location != "" {
		params.Add("geoUrn", location) // adjust with LinkedIn location URN if needed
	}
	if company != "" {
		params.Add("currentCompany", company)
	}
	if keywords != "" {
		params.Add("keywords", keywords)
	}

	return base + params.Encode()
}
