package entities

import (
	"bowlingGame/constants"
	"bowlingGame/scoring_strategy"
	"fmt"
)

type Player struct {
	ID              int
	Name            string
	TotalScore      int
	Scores          map[int]map[int]int // scores[chance][turn]
	ScoringStrategy scoring_strategy.ScoringStrategy
}

func NewPlayer(id int, name string) *Player {
	return &Player{
		ID:              id,
		Name:            name,
		Scores:          make(map[int]map[int]int),
		ScoringStrategy: &scoring_strategy.DefaultScoring{},
	}
}

func (p *Player) UpdateScoringStrategy(chance int) {
	if p.IsPreviousStrike(chance) {
		fmt.Println("doubling", p.Name)
		p.ScoringStrategy = &scoring_strategy.StrikeScoring{}
	} else {
		p.ScoringStrategy = &scoring_strategy.DefaultScoring{}
	}

}

// SetScore records the score for a specific chance and turn, adjusting for strikes.
func (p *Player) SetScore(chance, turn, score int) {
	if score < 0 {
		// Handle invalid score input
		return
	}

	//if p.IsPreviousStrike(chance) {
	//	fmt.Printf("%s Spare!!!!!!\n", p.Name)
	//	score *= 2 // Double the score if the previous chance was a strike
	//}
	p.UpdateScoringStrategy(chance)
	score = p.ScoringStrategy.CalculateScore(score)
	if p.Scores[chance] == nil {
		p.Scores[chance] = make(map[int]int)
	}
	p.Scores[chance][turn] = score
	p.TotalScore += score
}

// IsPreviousStrike checks if the previous chance was a strike.
func (p *Player) IsPreviousStrike(chance int) bool {
	return chance > 0 && p.IsSpare(chance-1)
}

// IsSpare checks if the previous chance was a spare.
func (p *Player) IsSpare(chance int) bool {
	return p.GetScore(chance, 0)+p.GetScore(chance, 1) == constants.STRIKE
}

// // IsStrike checks if the current turn's score is a strike.
func (p *Player) IsStrike(turn, score int) bool {
	return turn == 0 && score == constants.STRIKE
}

// GetScore retrieves the score for a specific chance and turn.
func (p *Player) GetScore(chance, turn int) int {
	if p.Scores[chance] != nil {
		return p.Scores[chance][turn]
	}
	return 0 // Return 0 if no score exists
}

func (p *Player) GetTotalScore() int {
	return p.TotalScore
}

func (p *Player) GetID() int {
	return p.ID
}

func (p *Player) GetName() string {
	return p.Name
}
