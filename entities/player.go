package entities

import "bowlingGame/constants"

type Player struct {
	ID         int
	Name       string
	TotalScore int
	Scores     map[int]map[int]int // scores[chance][turn]
}

func NewPlayer(id int, name string) *Player {
	return &Player{
		ID:     id,
		Name:   name,
		Scores: make(map[int]map[int]int),
	}
}

// SetScore records the score for a specific chance and turn, adjusting for strikes.
func (p *Player) SetScore(chance, turn, score int) {
	if score < 0 {
		// Handle invalid score input
		return
	}

	if p.IsPreviousStrike(chance) {
		score *= 2 // Double the score if the previous chance was a strike
	}
	if p.Scores[chance] == nil {
		p.Scores[chance] = make(map[int]int)
	}
	p.Scores[chance][turn] = score
	p.TotalScore += score
}

// IsPreviousStrike checks if the previous chance was a strike.
func (p *Player) IsPreviousStrike(chance int) bool {
	return chance > 0 && p.IsStrike(chance-1, p.GetScore(chance-1, 0))
}

// IsSpare checks if the previous chance was a spare.
func (p *Player) IsSpare(chance int) bool {
	return p.GetScore(chance, 0)+p.GetScore(chance, 1) == constants.STRIKE
}

// IsStrike checks if the current turn's score is a strike.
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
