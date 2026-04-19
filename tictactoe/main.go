package main

import (
	"fmt"
	"math/rand/v2"
)

type Symbol string

const (
	X_Symbol     Symbol = "X"
	O_Symbol     Symbol = "O"
	Empty_Symbol Symbol = "."
)

func getCount(symbol Symbol) int {
	switch symbol {
	case X_Symbol:
		return 1
	case O_Symbol:
		return -1
	default:
		return 0
	}
}

type Move struct {
	row int
	col int
}

type MoveStrategy interface {
	nextMove(b *Board) Move
}

type RandomMoveStrategy struct{}

func (r RandomMoveStrategy) nextMove(b *Board) Move {
	availableMoves := b.getAvailableMoves()
	size := len(availableMoves)

	randomNumber := rand.IntN(size)
	return Move{availableMoves[randomNumber][0], availableMoves[randomNumber][1]}
}

type Player struct {
	name         string
	symbol       Symbol
	moveStrategy MoveStrategy
}

func NewPlayer(name string, symbol Symbol, moveStrategy MoveStrategy) *Player {
	return &Player{
		name:         name,
		symbol:       symbol,
		moveStrategy: moveStrategy,
	}
}

func (p *Player) getSymbol() Symbol {
	return p.symbol
}

func (p *Player) nextMove(board *Board) Move {
	return p.moveStrategy.nextMove(board)
}

type Board struct {
	grid      [][]Symbol
	size      int
	movesLeft int
}

func newBoard(size int) *Board {
	grid := make([][]Symbol, size)
	for i := range grid {
		grid[i] = make([]Symbol, size)
		for j := range grid[i] {
			grid[i][j] = Empty_Symbol
		}
	}

	return &Board{
		grid:      grid,
		size:      size,
		movesLeft: size * size,
	}
}

func (b *Board) isFull() bool {
	return b.movesLeft == 0
}

func (b *Board) getSize() int {
	return b.size
}

func (b *Board) getCell(row int, col int) Symbol {
	return b.grid[row][col]
}

func (b *Board) placeMove(x int, y int, symbol Symbol) bool {
	if x < 0 || x >= b.size || y < 0 || y >= b.size {
		return false
	}
	if b.grid[x][y] != Empty_Symbol {
		return false
	}
	b.grid[x][y] = symbol
	b.movesLeft--

	return true
}

func (b *Board) getAvailableMoves() [][]int {
	moves := [][]int{}
	for i := range b.grid {
		for j := range b.grid[i] {
			if b.grid[i][j] == Empty_Symbol {
				moves = append(moves, []int{i, j})
			}
		}
	}
	return moves
}

type WinChecker struct {
	size      int
	cols      []int
	rows      []int
	diags     int
	antiDiags int
}

func newWinChecker(n int) *WinChecker {
	return &WinChecker{
		size:      n,
		cols:      make([]int, n),
		rows:      make([]int, n),
		diags:     0,
		antiDiags: 0,
	}
}
func (w *WinChecker) recordMove(row, col int, symbol Symbol) bool {
	count := getCount(symbol)

	w.rows[row] += count
	w.cols[col] += count
	if row == col {
		w.diags += count
	}
	if row+col == w.size-1 {
		w.antiDiags += count
	}
	target := w.size * count

	return w.rows[row] == target || w.cols[col] == target || w.diags == target || w.antiDiags == target
}

type GameState string

const (
	IN_PROGRESS GameState = "IN_PROGRESS"
	X_WON       GameState = "X_WON"
	Y_WON       GameState = "Y_WON"
	DRAW        GameState = "DRAW"
	OVER        GameState = "OVER"
)

type Game struct {
	board       *Board
	currentTurn int
	players     []Player
	winChecker  *WinChecker
	gameState   GameState
}

func newGame(size int, players []Player) *Game {
	return &Game{
		board:       newBoard(size),
		currentTurn: 0,
		players:     players,
		winChecker:  newWinChecker(size),
		gameState:   IN_PROGRESS,
	}
}

func (g *Game) getGameState() GameState {
	return g.gameState
}

func (g *Game) makeMove(row, col int) GameState {
	if g.gameState != IN_PROGRESS {
		return OVER
	}

	currentPlayer := g.players[g.currentTurn]

	if !g.board.placeMove(row, col, currentPlayer.symbol) {
		fmt.Println("Invalid move to row & col ", row, col)
		return IN_PROGRESS
	}

	if g.winChecker.recordMove(row, col, currentPlayer.symbol) {
		if currentPlayer.symbol == X_Symbol {
			g.gameState = X_WON
		} else {
			g.gameState = Y_WON
		}
	} else if g.board.isFull() {
		g.gameState = DRAW
	} else {
		g.currentTurn = (g.currentTurn + 1) % len(g.players)
	}

	return g.gameState
}

func (g *Game) play() {
	fmt.Println("starting the game")
	for g.gameState == IN_PROGRESS {
		player := g.players[g.currentTurn]
		nextMove := player.nextMove(g.board)
		fmt.Println("next move by player ", player.symbol, "to row and col ", nextMove.row, nextMove.col)
		g.makeMove(nextMove.row, nextMove.col)
	}

	fmt.Println("game state final ", g.gameState)
}

func main() {
	players := []Player{
		{
			name:         "x",
			symbol:       X_Symbol,
			moveStrategy: RandomMoveStrategy{},
		},
		{
			name:         "y",
			symbol:       O_Symbol,
			moveStrategy: RandomMoveStrategy{},
		},
	}
	game := newGame(3, players)

	game.play()
}
