package stealth

import "github.com/go-rod/rod"

// ApplyFingerprint injects JS BEFORE any site loads
func ApplyFingerprint(page *rod.Page) {
	page.MustEvalOnNewDocument(`
		Object.defineProperty(navigator, 'webdriver', {
			get: () => undefined
		});

		Object.defineProperty(navigator, 'languages', {
			get: () => ['en-US', 'en']
		});

		Object.defineProperty(navigator, 'platform', {
			get: () => 'Win32'
		});
	`)
}
