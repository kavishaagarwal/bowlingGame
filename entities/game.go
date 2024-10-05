package entities

import (
	"container/list"
	"fmt"
	"sort"
)

type Game struct {
	Lane       *Lane
	Ball       *Ball
	Players    *list.List
	TotalTurns int
}

func NewGame(names []string, TotalTurns int) *Game {
	g := &Game{}
	g.initializeGame(names, TotalTurns)
	return g
}

func (g *Game) initializeGame(names []string, TotalTurns int) {
	g.Ball = NewBowl()
	g.addPlayers(names)
	g.Lane = NewLane(1)
	g.TotalTurns = TotalTurns
}

func (g *Game) addPlayers(names []string) {
	g.Players = list.New()
	for idx, name := range names {
		player := NewPlayer(idx+1, name)
		g.Players.PushBack(player)
	}
}

func (g *Game) findPlayerTurn() *Player {
	player := g.Players.Remove(g.Players.Front())
	g.Players.PushBack(player)
	return player.(*Player)
}

func (g *Game) GetScores() {
	for e := g.Players.Front(); e != nil; e = e.Next() {
		player := e.Value.(*Player)
		fmt.Print(fmt.Sprintf("Player: %s Score: %d *** ", player.Name, player.GetTotalScore()))
	}
	fmt.Println()
}

func (g *Game) StartGame() {
	g.Play()
	g.DeclareWinner()
}

func (g *Game) Play() {
	for i := 0; i < g.TotalTurns; i++ {
		for p := g.Players.Front(); p != nil; p = p.Next() {
			player := p.Value.(*Player)
			for j := 0; j < 2; j++ {
				lastScore := 0
				if j == 1 {
					lastScore = player.Scores[i][0]
				}
				score := g.Ball.RollBowl(lastScore)
				fmt.Println(player.Name, score)
				player.SetScore(i, j, score)
				g.GetScores()
				if player.IsStrike(j, score) {
					fmt.Println(fmt.Sprintf("%s Strike!!!!!!", player.Name))
					break
				}
			}
		}
	}
}

func (g *Game) DeclareWinner() {
	var playersSlice []*Player
	for e := g.Players.Front(); e != nil; e = e.Next() {
		player := e.Value.(*Player)
		playersSlice = append(playersSlice, player)
	}
	sort.Slice(playersSlice, func(i, j int) bool {
		return playersSlice[i].TotalScore > playersSlice[j].TotalScore
	})
	fmt.Println(fmt.Sprintf("Winner is %d %s and total score is %d", playersSlice[0].GetID(), playersSlice[0].GetName(), playersSlice[0].GetTotalScore()))
}
