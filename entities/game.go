package entities

import (
	"fmt"
	"sort"
)

const MaxRollsPerTurn = 2

type Game struct {
	Ball       *Ball
	Players    []*Player
	TotalTurns int
	LaneID     int // Track which lane this game is using
}

// NewGame initializes a new game for a specific lane.
func NewGame(names []string, totalTurns int, laneID int) *Game {
	g := &Game{
		LaneID: laneID,
	}
	g.initializeGame(names, totalTurns)
	return g
}

// Update the initializeGame method to handle lanes if needed.
func (g *Game) initializeGame(names []string, totalTurns int) {
	g.Ball = NewBall()
	g.addPlayers(names)
	g.TotalTurns = totalTurns
}

func (g *Game) addPlayers(names []string) {
	g.Players = make([]*Player, len(names))
	for idx, name := range names {
		g.Players[idx] = NewPlayer(idx+1, name)
	}
}

func (g *Game) StartGame() {
	fmt.Println("\nStarting new game!!!!!!!!!!!!!!")
	g.Play()
	g.DeclareWinner()
}

func (g *Game) Play() {
	for i := 0; i < g.TotalTurns; i++ {
		for _, player := range g.Players {
			for j := 0; j < MaxRollsPerTurn; j++ {
				lastScore := 0
				if j == 1 {
					lastScore = player.Scores[i][0] // Last score for second roll
				}
				score := g.Ball.Roll(lastScore)
				fmt.Print(player.Name, " ", score, "             ")
				player.SetScore(i, j, score)
				if player.IsStrike(j, score) {
					fmt.Printf("%s Strike!!!!!!\n", player.Name)
					break
				}
			}
		}
		fmt.Println()
	}
}

func (g *Game) DeclareWinner() {
	sort.Slice(g.Players, func(i, j int) bool {
		return g.Players[i].TotalScore > g.Players[j].TotalScore
	})
	winner := g.Players[0]
	fmt.Println()
	fmt.Printf("\nWinner is %s and total score is %d!!!!!!!!!!!!!!!!!!!!!\n", winner.GetName(), winner.GetTotalScore())
}
