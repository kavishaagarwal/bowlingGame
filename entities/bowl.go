package entities

import (
	"math/rand"
	"time"
)

type Ball struct {
	//Number int
}

func NewBowl() *Ball {
	return &Ball{
		//Number: number,
	}
}

func (b *Ball) RollBowl(lastScore int) int {
	rand.Seed(time.Now().UnixNano())
	//Pins dropped can be fromm 0 to 10
	if lastScore > 10 {
		lastScore = lastScore / 2
	}
	pinsDropped := rand.Intn(11 - lastScore)
	return pinsDropped
}
