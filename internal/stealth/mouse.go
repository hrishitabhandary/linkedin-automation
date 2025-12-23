package stealth

import (
	"math/rand"

	"github.com/go-rod/rod"
)

// ClickHuman simulates a human-like click using hover + delay
func ClickHuman(el *rod.Element) {
	// Hover before clicking (human behavior)
	_ = el.Hover()

	// Small reaction delay
	React(120, 280)

	// Click
	el.MustClick()

	// Post-click pause
	React(80, 160)
}

// HoverHuman simulates human hovering
func HoverHuman(el *rod.Element) {
	_ = el.Hover()
	React(200, 400)
}

// RandomIdleMouse simulates cursor inactivity
func RandomIdleMouse() {
	Idle(300, 700)
	_ = rand.Intn(10) // entropy placeholder
}
