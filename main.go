package main

import (
	"bowlingGame/admin"
	"bowlingGame/entities"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	arena := entities.NewArena(3)
	admin := admin.NewAdmin(arena)
	admin.AddLanes(1)
	ss, _ := json.Marshal(arena)
	fmt.Println(string(ss))
	// Start multiple games on different lanes
	go arena.StartGameOnLane([]string{"Kavu", "Bittu"}, 10)
	go arena.StartGameOnLane([]string{"Golu", "Akash"}, 10)
	go arena.StartGameOnLane([]string{"Golu", "Akash"}, 10)
	go arena.StartGameOnLane([]string{"Golu", "Akash"}, 10)
	go arena.StartGameOnLane([]string{"Golu", "Akash"}, 10)
	go arena.StartGameOnLane([]string{"Golu", "Akash"}, 10)

	// You can start more games on different lanes as needed
	//arena.StartGameOnLane(3, []string{"Alice", "Bob"}, 10)

	// Display the state of each game if needed
	for laneID, game := range arena.Games {
		ss, err := json.Marshal(game)
		if err != nil {
			log.Fatalf("Error marshaling game on lane %d to JSON: %v", laneID, err)
		}
		fmt.Printf("Game on Lane %d: %s\n", laneID, string(ss))
	}
	time.Sleep(3 * time.Second)
}
