package browser

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// StartBrowser launches Chromium and returns a connected browser instance
func StartBrowser() *rod.Browser {
	u := launcher.New().
		Headless(false).   // visible browser (important for LinkedIn)
		Leakless(false).   // avoids Windows antivirus issues
		MustLaunch()

	browser := rod.New().
		ControlURL(u).
		MustConnect()

	log.Println("✅ Browser launched and connected")
	return browser
}
