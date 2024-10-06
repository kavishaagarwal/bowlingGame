package scoring_strategy

import "fmt"

type ScoringStrategy interface {
	CalculateScore(currentScore int) int
}

type DefaultScoring struct{}

func (d *DefaultScoring) CalculateScore(currentScore int) int {
	return currentScore
}

type StrikeScoring struct{}

func (s *StrikeScoring) CalculateScore(currentScore int) int {
	fmt.Println("Score doubled!!!!!")
	return 2 * currentScore
}
