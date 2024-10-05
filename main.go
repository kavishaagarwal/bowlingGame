package main

import (
	"bowlingGame/entities"
	"encoding/json"
	"fmt"
)

func main() {
	//arena := entities.NewArena()
	//admin := admin.NewAdmin(arena)
	//admin.AddLanes(3)

	playerNames := []string{"Kavu", "Bittu", "Golu", "Akash"}
	game := entities.NewGame(playerNames, 10)
	ss, _ := json.Marshal(game)
	fmt.Println("game:", string(ss))
	//display := display.NewDisplayBoard(game)
	//display.GetScores()

	game.StartGame()

	//display.GetScores()
}
