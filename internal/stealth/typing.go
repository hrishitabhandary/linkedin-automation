package stealth

import (
	"math/rand"

	"github.com/go-rod/rod"
)

func HumanType(el *rod.Element, text string) error {
	for i, ch := range text {

		// Keystroke delay
		React(70, 200)

		if err := el.Input(string(ch)); err != nil {
			return err
		}

		// Occasional typo + correction
		if rand.Float64() < 0.06 && i > 2 {
			React(80, 140)
			_ = el.Input("x")
			React(80, 120)
			_ = el.Input("\b") // Backspace
		}

		// Thinking pause
		if rand.Float64() < 0.04 {
			Think(300, 700)
		}
	}
	return nil
}
