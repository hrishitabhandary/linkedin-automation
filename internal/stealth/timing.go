package stealth

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Think simulates cognitive delay
func Think(minMs, maxMs int) {
	sleep(minMs, maxMs)
}

// React simulates reaction time before an action
func React(minMs, maxMs int) {
	sleep(minMs, maxMs)
}

// Idle simulates longer idle pauses
func Idle(minMs, maxMs int) {
	sleep(minMs, maxMs)
}

func sleep(minMs, maxMs int) {
	if maxMs <= minMs {
		time.Sleep(time.Duration(minMs) * time.Millisecond)
		return
	}
	time.Sleep(time.Duration(minMs+rand.Intn(maxMs-minMs)) * time.Millisecond)
}
