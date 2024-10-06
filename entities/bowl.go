package entities

import (
	"math/rand"
	"time"
)

const MaxPins = 10

type Ball struct{}

func init() {
	rand.Seed(time.Now().UnixNano()) // Initialize random seed once
}

func NewBall() *Ball {
	return &Ball{}
}

// Roll simulates rolling the ball and returning the number of pins dropped.
func (b *Ball) Roll(lastScore int) int {
	if lastScore < 0 || lastScore > MaxPins {
		// Handle invalid lastScore
		return 0
	}

	// Calculate available pins based on lastScore
	pinsAvailable := MaxPins - lastScore
	if pinsAvailable < 0 {
		pinsAvailable = 0
	}

	pinsDropped := rand.Intn(pinsAvailable + 1) // +1 to include pinsAvailable
	return pinsDropped
}
