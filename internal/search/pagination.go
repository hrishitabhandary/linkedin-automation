package search

import (
	"time"

	"github.com/go-rod/rod"
)

// NextPages clicks through a given number of search result pages
func NextPages(page *rod.Page, pages int) {
	for i := 0; i < pages; i++ {
		nextBtn := page.MustElement("button[aria-label='Next']")
		nextBtn.MustClick()
		page.MustWaitLoad()
		time.Sleep(2 * time.Second) // allow content to render
	}
}
