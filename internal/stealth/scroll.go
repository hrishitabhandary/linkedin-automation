package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// Scroll to top smoothly
func ScrollToTop(page *rod.Page) {
	page.MustEval(`() => {
		window.scrollTo({ top: 0, behavior: "smooth" })
	}`)
	time.Sleep(2 * time.Second)
}
// HumanScroll scrolls the page like a human over the given duration
func HumanScroll(page *rod.Page, duration time.Duration) {
	end := time.Now().Add(duration)

	for time.Now().Before(end) {
		// Random scroll distance
		y := float64(rand.Intn(400) + 100)
		page.Mouse.Scroll(0, y, 5) // 5 steps for smoother scroll
		time.Sleep(time.Duration(rand.Intn(300)+200) * time.Millisecond)
	}
}
