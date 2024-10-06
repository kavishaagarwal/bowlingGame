package display

import (
	"bowlingGame/entities"
	"fmt"
)

type GameView struct {
	Game *entities.Game
}

func NewGameView(game *entities.Game) *GameView {
	return &GameView{Game: game}
}

func (gv *GameView) DisplayScores() {
	for _, player := range gv.Game.Players {
		fmt.Printf("Player: %s Score: %d *** ", player.Name, player.GetTotalScore())
	}
	fmt.Println()
}
