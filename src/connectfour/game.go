package connectfour

import (
	"connectfour/ai"
	"connectfour/board"
	"math/rand"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Game struct {
	Board         board.Board
	CurrentPlayer int
	IsFinished    bool
	Winner        int
}

const HumanPlayer int = 1
const AIPlayer int = 2

func New() Game {
	return Game{
		Board:         board.Board{},
		CurrentPlayer: rand.Intn(1) + 1,
	}
}

func (game *Game) HumanPlay(column int) (int, error) {
	nextBoard, err := game.Board.AddDisc(column, HumanPlayer)
	game.Board = nextBoard
	game.checkGameStatus()
	game.switchPlayer()
	return column, err
}

func (game *Game) ComputerPlay() (int, error) {
	column, err := ai.NextBestMoveInTime(game.Board, AIPlayer, time.Second)
	if err != nil {
		return column, err
	}
	nextBoard, err := game.Board.AddDisc(column, AIPlayer)
	if err != nil {
		return column, err
	}
	game.Board = nextBoard
	check(err)
	game.checkGameStatus()
	game.switchPlayer()
	return column, nil
}

func (game *Game) switchPlayer() {
	if game.CurrentPlayer == HumanPlayer {
		game.CurrentPlayer = AIPlayer
	} else {
		game.CurrentPlayer = HumanPlayer
	}
}

func (game *Game) checkGameStatus() {
	currentPlayerHasWon := ai.HasWon(game.Board, game.CurrentPlayer)
	boardIsFull := game.Board.IsFull()

	if currentPlayerHasWon {
		game.IsFinished = true
		game.Winner = game.CurrentPlayer
	} else if boardIsFull {
		game.IsFinished = true
	}
}
