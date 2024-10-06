package admin

import (
	"bowlingGame/entities"
	"fmt"
)

// go mod init github.com/kavishagarwal/bowlingGame
type Admin struct {
	Arena *entities.Arena
}

func NewAdmin(arena *entities.Arena) *Admin {
	return &Admin{
		Arena: arena,
	}
}

func (a *Admin) AddLanes(number int) {
	for i := 0; i < number; i++ {
		a.Arena.AddLane()
	}
}

func (a *Admin) RemoveLanes(id int) {
	if id >= len(a.Arena.GetLanes()) {
		return
	}
	if a.Arena.GetLanes()[id].IsOccupied() {
		fmt.Println("Game is on, can't remove the lane")
		return
	}
	a.Arena.RemoveLane(id)
}
