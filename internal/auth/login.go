// package auth

// import (
// 	"log"
// 	"math/rand"
// 	"time"

// 	"github.com/go-rod/rod"
// )

// // humanType types text slowly like a real human
// func humanType(el *rod.Element, text string) {
// 	for _, ch := range text {
// 		el.MustInput(string(ch))
// 		// random delay between keystrokes (80–200ms)
// 		time.Sleep(time.Duration(80+rand.Intn(120)) * time.Millisecond)
// 	}
// }

// // PerformLogin logs into LinkedIn using human-like behavior
// func PerformLogin(page *rod.Page, email, password string) error {
// 	log.Println("🔐 Attempting LinkedIn login (human-like)")

// 	rand.Seed(time.Now().UnixNano())

// 	// Email field
// 	emailEl := page.MustElement(`#username`)
// 	emailEl.MustClick()
// 	humanType(emailEl, email)

// 	time.Sleep(1 * time.Second)

// 	// Password field
// 	passEl := page.MustElement(`#password`)
// 	passEl.MustClick()
// 	humanType(passEl, password)

// 	time.Sleep(1 * time.Second)

// 	// Click Sign in
// 	page.MustElement(`button[type="submit"]`).MustClick()

// 	// Wait for navigation
// 	page.MustWaitLoad()
// 	time.Sleep(5 * time.Second)

// 	log.Println("✅ Login attempt finished")
// 	return nil
// }
 

package auth

import (
	"encoding/json"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// humanType types text slowly like a real human
func humanType(el *rod.Element, text string) {
	for _, ch := range text {
		el.MustInput(string(ch))
		time.Sleep(time.Duration(80+rand.Intn(120)) * time.Millisecond)
	}
}

// PerformLogin performs human-like login
func PerformLogin(page *rod.Page, email, password string) {
	rand.Seed(time.Now().UnixNano())

	log.Println("🔐 Performing human-like login")

	// Email field
	emailEl := page.MustElement(`#username`)
	emailEl.MustClick()
	humanType(emailEl, email)
	time.Sleep(1 * time.Second)

	// Password field
	passEl := page.MustElement(`#password`)
	passEl.MustClick()
	humanType(passEl, password)
	time.Sleep(1 * time.Second)

	// Click Sign in
	page.MustElement(`button[type="submit"]`).MustClick()
	// Wait for response
	page.MustWaitLoad()
	time.Sleep(5 * time.Second)

	// 🔍 Detect security checkpoints
   if checkpoint := DetectSecurityCheckpoint(page); checkpoint != "" {
	log.Println("⚠️ Security checkpoint:", checkpoint)
	log.Println("🛑 Manual intervention required")
	time.Sleep(60 * time.Second) // keep browser open for user
   return
   }

	log.Println("✅ Login finished")
}

// MustSaveCookies saves all cookies of the page to a file
func MustSaveCookies(page *rod.Page, filePath string) {
	cookies, err := page.Cookies([]string{})
	if err != nil {
		log.Fatal("Failed to get cookies:", err)
	}

	data, err := json.Marshal(cookies)
	if err != nil {
		log.Fatal("Failed to marshal cookies:", err)
	}

	// Ensure directory exists
	os.MkdirAll("data", os.ModePerm)

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		log.Fatal("Failed to save cookies:", err)
	}
}

// MustReadCookies reads cookies from a file and converts them to []*proto.NetworkCookieParam
func MustReadCookies(filePath string) []*proto.NetworkCookieParam {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal("Failed to read cookies:", err)
	}

	var cookies []*proto.NetworkCookie
	if err := json.Unmarshal(data, &cookies); err != nil {
		log.Fatal("Failed to unmarshal cookies:", err)
	}

	var cookieParams []*proto.NetworkCookieParam
	for _, c := range cookies {
		cookieParams = append(cookieParams, &proto.NetworkCookieParam{
			Name:     c.Name,
			Value:    c.Value,
			Domain:   c.Domain,
			Path:     c.Path,
			Expires:  c.Expires,
			HTTPOnly: c.HTTPOnly,
			Secure:   c.Secure,
			SameSite: c.SameSite,
		})
	}

	return cookieParams
}

// LoginWithCookies handles automatic login using cookies if available
func LoginWithCookies(page *rod.Page, email, password string) {
	cookieFile := "data/cookies.json"

	// ✅ Try loading cookies
	if _, err := os.Stat(cookieFile); err == nil {
		page.MustSetCookies(MustReadCookies(cookieFile)...)
		log.Println("🍪 Cookies loaded, bypassing login")
		page.MustNavigate("https://www.linkedin.com/feed/")
		page.MustWaitLoad()
		return
	}

	// 🔐 No cookies → human-like login
	log.Println("🔐 No cookies found, performing login")
	PerformLogin(page, email, password)

	// 🍪 Save cookies for next time
	MustSaveCookies(page, cookieFile)
	log.Println("🍪 Cookies saved for future logins")
}


// DetectSecurityCheckpoint checks for captcha, 2FA, or account chooser
func DetectSecurityCheckpoint(page *rod.Page) string {
	// CAPTCHA
	if _, err := page.Element(`iframe[src*="captcha"]`); err == nil {
		return "CAPTCHA detected"
	}

	// 2FA OTP
	if _, err := page.Element(`input[name="pin"]`); err == nil {
		return "2FA detected"
	}

	// Account chooser / resume session
	if page.MustHasR("body", "Welcome back") {
		return "Account chooser detected"
	}

	return ""
}
