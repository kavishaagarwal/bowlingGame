package entities

import (
	"fmt"
	"sync"
)

type Arena struct {
	Lanes []*Lane
	Games map[int]*Game // Map of lane IDs to Games
	mu    sync.Mutex
}

// NewArena creates a new arena with the specified number of lanes.
func NewArena(numLanes int) *Arena {
	arena := &Arena{
		Games: make(map[int]*Game),
	}
	for i := 1; i <= numLanes; i++ {
		arena.AddLane()
	}
	return arena
}

func (a *Arena) StartGameOnLane(playerNames []string, totalTurns int) {
	a.mu.Lock()         // Lock the mutex before modifying shared state
	defer a.mu.Unlock() // Ensure the mutex is unlocked after the function returns

	lane := a.GetAvailableLanes()
	if lane == nil {
		fmt.Println("No lanes available")
	}

	lane.SetOccupied(true)
	game := NewGame(playerNames, totalTurns, lane.ID)
	a.SetGameOnLaneId(lane.ID, game)
	game.StartGame()
	a.FreeLane(lane.ID)
}

func (a *Arena) AddLane() {
	l := NewLane(len(a.Lanes) + 1)
	a.Lanes = append(a.Lanes, l)
}

func (a *Arena) GetLanes() []*Lane {
	return a.Lanes
}

func (a *Arena) GetAvailableLanes() *Lane {
	for _, lane := range a.Lanes {
		if !lane.isOccupied {
			return lane
		}
	}
	return nil
}

func (a *Arena) RemoveLane(laneNumber int) {
	a.Lanes = append(a.Lanes[:laneNumber], a.Lanes[laneNumber+1:]...)
}

func (a *Arena) SetGameOnLaneId(laneNumber int, game *Game) {
	a.Games[laneNumber] = game
}

func (a *Arena) FreeLane(laneNumber int) {
	delete(a.Games, laneNumber)
	lane := a.Lanes[laneNumber-1]
	lane.SetOccupied(false)
	fmt.Println(fmt.Sprintf("Lane no %d is free now", laneNumber))
}
