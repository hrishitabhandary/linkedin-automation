package connect

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

// ClickConnect clicks the "Connect" button on a LinkedIn profile.
func ClickConnect(p *rod.Page) bool {
	btn := p.MustElementR("button", "Connect")
	if btn == nil {
		return false
	}

	// Scroll into view and click
	btn.ScrollIntoView()
	btn.Click(proto.InputMouseButtonLeft, 1)
	return true
}

// SendNote sends a personalized note when connecting
func SendNote(p *rod.Page, note string) {
	// Wait for the note textarea to appear
	textarea := p.MustElement("textarea[name=message]")
	if textarea == nil {
		log.Println("No note textarea found")
		return
	}

	// Type the note like a human
	textarea.MustFocus()
	textarea.Input(note)

	// Click the "Send" button
	sendBtn := p.MustElementR("button", "Send")
	if sendBtn != nil {
		sendBtn.Click(proto.InputMouseButtonLeft, 1)
	}
}

// Tracker keeps track of sent connection requests
type Tracker struct {
	Seen map[string]bool
}

// NewTracker initializes a new Tracker
func NewTracker() *Tracker {
	return &Tracker{
		Seen: make(map[string]bool),
	}
}

// CanSend checks if more requests can be sent
func (t *Tracker) CanSend() bool {
	// Example daily limit: 10
	return len(t.Seen) < 10
}

// MarkSent marks a profile URL as already sent
func (t *Tracker) MarkSent(url string) {
	t.Seen[url] = true
}
