package main

import (
	
	
	"log"
	
	
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	

	"github.com/hrishitabhandary/linkedin-automation-go/internal/auth"
	"github.com/hrishitabhandary/linkedin-automation-go/internal/config"
	"github.com/hrishitabhandary/linkedin-automation-go/internal/stealth"
	"github.com/hrishitabhandary/linkedin-automation-go/internal/search"
)


	

func main() {


	// 🔐 Load LinkedIn credentials
	email, password, err := config.LoadCredentials()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Loaded email from ENV:", email)

	// 🌐 Launch browser
	url := launcher.New().
		Headless(false).
		Leakless(false).
		Set("start-maximized", "true").
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	defer browser.MustClose()

	page := browser.MustPage()
	stealth.ApplyFingerprint(page)

	// 🔑 Login
	page.MustNavigate("https://www.linkedin.com/login")
	page.MustWaitLoad()
	auth.LoginWithCookies(page, email, password)

	page.MustElementR("div", "Start a post")
	time.Sleep(3 * time.Second)

	// 🔗 SEND CONNECTION REQUESTS
	connectFromNetwork(page, 3, 5)

	// ---------------- SEARCH ----------------
	searchBox := page.MustElement("input[placeholder='Search']")
	searchBox.MustFocus()
	_ = stealth.HumanType(searchBox, "Golang Developer")
	_ = searchBox.Input("\n")
	time.Sleep(2 * time.Second)

	searchQuery := "Golang Developer"
	searchURL := search.BuildSearchURL(searchQuery, "India", "", "")
	page.MustNavigate(searchURL)
	page.MustWaitLoad()
	time.Sleep(2 * time.Second)

	// 📩 MESSAGE FIRST 3 SEARCH SUGGESTIONS
    messageTopSuggestions(page)

	// ---------------- COLLECT PROFILES ----------------
	profileURLs := []string{}
	for pageNum := 0; pageNum < 3; pageNum++ {
		stealth.HumanScroll(page, 5*time.Second)
		time.Sleep(2 * time.Second)

		results := page.MustElements("div.entity-result__item")
		log.Println("Found profiles on page", pageNum+1, ":", len(results))

		for _, r := range results {
			linkEl := r.MustElement("a.app-aware-link")
			href, _ := linkEl.Attribute("href")
			if href != nil {
				profileURLs = append(profileURLs, *href)
			}
		}

		nextBtn := page.MustElementR("button", "Next")
		if nextBtn == nil {
			break
		}
		stealth.ClickHuman(nextBtn)
		page.MustWaitLoad()
		time.Sleep(3 * time.Second)
	}

	uniqueProfiles := search.RemoveDuplicates(profileURLs)
	log.Println("Unique profiles collected:", len(uniqueProfiles))
	for _, u := range uniqueProfiles {
		log.Println(u)
	}


	
}


func connectFromNetwork(page *rod.Page, min, max int) {
	count := 0
	limit := min + time.Now().Nanosecond()%(max-min+1)

	log.Println("Sending", limit, "connection requests")

	page.MustNavigate("https://www.linkedin.com/mynetwork/")
	page.MustWaitLoad()
	time.Sleep(5 * time.Second)

	for count < limit {
		// Find Connect buttons
		connectBtns := page.MustElementsX("//button[.//span[text()='Connect']]")

		if len(connectBtns) == 0 {
			log.Println("No connect buttons, scrolling...")
			page.Mouse.Scroll(0, 1200, 5)

			time.Sleep(3 * time.Second)
			continue
		}

		for _, btn := range connectBtns {
			if count >= limit {
				break
			}

			btn.MustScrollIntoView()
			time.Sleep(1 * time.Second)
			btn.MustClick()
			time.Sleep(2 * time.Second)

			// Handle modal
			sendBtn, err := page.Timeout(3 * time.Second).
				ElementX("//button[.//span[text()='Send without a note']]")
			if err == nil {
				sendBtn.MustClick()
			}

			count++
			log.Println("✅ Connection sent:", count)
			time.Sleep(4 * time.Second)
		}
	}

	log.Println("🎯 Finished sending requests")
}

func messageTopSuggestions(page *rod.Page) {
	log.Println("📩 Messaging top 3 suggestions")

	// Wait for search results
	page.MustWaitLoad()
	time.Sleep(4 * time.Second)

	// Get all visible Message buttons
	messageBtns := page.MustElementsX(
		"//button[.//span[text()='Message']]",
	)

	if len(messageBtns) == 0 {
		log.Println("❌ No Message buttons found")
		return
	}

	limit := 3
	if len(messageBtns) < limit {
		limit = len(messageBtns)
	}

	for i := 0; i < limit; i++ {
		btn := messageBtns[i]

		btn.MustScrollIntoView()
		time.Sleep(1 * time.Second)
		btn.MustClick()

		// 👇 ADD THE PREMIUM CHECK HERE (BEFORE msgBox)

        time.Sleep(2 * time.Second)

       // 🔒 Premium / InMail popup check
       _, err := page.Timeout(2 * time.Second).
	  ElementR("h2", "Message .* with Premium")


    if err == nil {
	 log.Println("⚠️ Premium popup detected, skipping")

	 closeBtn := page.MustElement("button[aria-label='Dismiss']")
	 closeBtn.MustClick()
	 time.Sleep(2 * time.Second)
	 continue
    }

        // ✅ ONLY AFTER THIS, wait for message box
		// Wait for message box
		msgBox := page.MustElement(
			"div[contenteditable='true']",
		)

		message := "Hi! I came across your profile while searching for Golang developers. Would love to connect and exchange insights."

		_ = stealth.HumanType(msgBox, message)
		time.Sleep(1 * time.Second)

		// Click Send
		sendBtn := page.MustElementX(
			"//button[.//span[text()='Send']]",
		)
		sendBtn.MustClick()

		log.Println("✅ Message sent to suggestion", i+1)

		time.Sleep(4 * time.Second)

		// Close message window
		closeBtn, err := page.Timeout(2 * time.Second).
			Element("button.msg-overlay-bubble-header__control")
		if err == nil {
			closeBtn.MustClick()
		}

		time.Sleep(3 * time.Second)
	}
}
