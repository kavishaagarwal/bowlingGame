package display

import "bowlingGame/entities"

type DisplayBoard struct {
	Game *entities.Game
}

func NewDisplayBoard(game *entities.Game) *DisplayBoard {
	return &DisplayBoard{
		Game: game,
	}
}

func (db *DisplayBoard) GetScores() {
	db.Game.GetScores()
}
