package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	name     string
	position int
}

type Board struct {
	size    int
	ladders map[int]int
	snakes  map[int]int
}

func CreateBoard(size int, ladders map[int]int, snakes map[int]int) *Board {
	return &Board{
		size:    size,
		ladders: ladders,
		snakes:  snakes,
	}
}

type GameManager struct {
	Players         []*Player
	GameBoard       *Board
	CurrentPlayer   int
	NumberOfPlayers int
	IsGameStarted   bool
}

func CreateGame(board *Board) *GameManager {
	return &GameManager{
		Players:         make([]*Player, 0),
		GameBoard:       board,
		CurrentPlayer:   0,
		NumberOfPlayers: 0,
		IsGameStarted:   false,
	}
}

func (g *GameManager) AddPlayer(name string) {
	if g.IsGameStarted {
		fmt.Println("game already started, cant add players")
		return
	}
	g.Players = append(g.Players, &Player{
		name:     name,
		position: 0,
	})
	fmt.Println("Added player ", name)
	g.NumberOfPlayers++
}

func (g *GameManager) PrintPlayerPositions() {
	for _, player := range g.Players {
		fmt.Print("Player name : ", player.name, " position : ", player.position)
	}
	fmt.Println(" ")
}

func (g *GameManager) Start() {
	if g.NumberOfPlayers < 2 {
		fmt.Println("can't start the game, less than 2 players")
		return
	}
	g.IsGameStarted = true
	g.CurrentPlayer = 0
	for {
		// start with rolling dice
		playerNo := g.CurrentPlayer
		player := g.Players[playerNo]
		position := player.position
		name := player.name
		dice := rand.Intn(6) + 1 // 1-6 random number

		// moving the current player
		newPostion := position + dice

		// if snake and ladder print that
		isSnakeBitten := false
		newPositionAfterSnakeBite, ok1 := g.GameBoard.snakes[newPostion]
		if ok1 {
			// found snake
			isSnakeBitten = true
		}

		isLadderClimbed := false
		newPositionAfterLadder, ok2 := g.GameBoard.ladders[newPostion]
		if ok2 {
			isLadderClimbed = true
		}

		if isSnakeBitten {
			fmt.Println(name, " rolled a ", dice, " -> bitten by snake from ", newPostion, " to ", newPositionAfterSnakeBite)
			newPostion = newPositionAfterSnakeBite
		} else if isLadderClimbed {
			fmt.Println(name, " rolled a ", dice, " -> climbed ladder from ", newPostion, " to ", newPositionAfterLadder)
			newPostion = newPositionAfterLadder
		} else {
			fmt.Println(name, " rolled a ", dice, " -> moved from ", position, " to ", newPostion)
		}

		// if 100, declare the winner
		if newPostion >= 100 {
			fmt.Println("Player won: ", name)
			break
		}

		g.Players[playerNo].position = newPostion
		g.CurrentPlayer = (playerNo + 1) % g.NumberOfPlayers
		// move to the next player
		time.Sleep(300 * time.Millisecond)
	}
}

func main() {
	ladders := map[int]int{
		5:  23,
		7:  29,
		14: 55,
		25: 77,
		36: 88,
		13: 97,
	}
	snakes := map[int]int{
		27: 6,
		36: 12,
		58: 24,
		91: 30,
		73: 53,
		99: 50,
	}
	board := CreateBoard(100, ladders, snakes)

	game := CreateGame(board)
	game.AddPlayer("Avinash")
	game.AddPlayer("Priyanshu")
	game.AddPlayer("Ayush")

	game.Start()
}
