package entities

import "bowlingGame/constants"

type Player struct {
	ID         int
	Name       string
	TotalScore int
	Scores     map[int]map[int]int
}

func NewPlayer(id int, name string) *Player {
	//s := make(map[int]int)
	score := make(map[int]map[int]int)
	return &Player{
		ID:     id,
		Name:   name,
		Scores: score,
	}
}

func (p *Player) SetScore(chance, turn, score int) {
	if p.IsPreviousStrike(chance) {
		score = 2 * score
	}
	if p.Scores[chance] == nil {
		p.Scores[chance] = make(map[int]int)
	}
	p.Scores[chance][turn] = score
	p.TotalScore += score
}

func (p *Player) IsPreviousStrike(chance int) bool {
	return chance > 0 && p.IsSpare(chance-1)
}

func (p *Player) IsSpare(chance int) bool {
	return p.Scores[chance][0]+p.Scores[chance][1] == constants.STRIKE
}

func (p *Player) IsStrike(turn, score int) bool {
	return turn == 0 && score == constants.STRIKE
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
