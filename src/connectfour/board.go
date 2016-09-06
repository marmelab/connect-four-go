package connectfour

import (
	"errors"
	"strings"
)

const boardWidth int = 7
const boardHeight int = 6

const FirstPlayerChar string = "x"
const SecondPlayerChar string = "o"

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

type Board [boardHeight][boardWidth]uint8

func New(input string) Board {
	lines := strings.Split(input, "\n")
	board := Board{}
	for lineIndex, line := range lines {
		for charIndex, char := range line {
			switch string(char) {
			case FirstPlayerChar:
				board[lineIndex][charIndex] = 1
			case SecondPlayerChar:
				board[lineIndex][charIndex] = 2
			}
		}
	}
	return board
}

func (board Board) AddDisc(column, player uint8) (Board, error) {
	dropped := false

	for i := boardHeight - 1; i >= 0; i-- {
		if board[i][column] == 0 {
			dropped = true
			board[i][column] = player
			break
		}
	}

	if !dropped {
		return board, errors.New("Illegal move")
	}

	return board, nil
}

func (board Board) score(player uint8) int {

	playerFourAligned, playerThreeAligned, playerTwoAligned := board.numberOfAlignedDiscs(player)

	opponent := getOpponent(player)
	_, _, opponentFourAligned := board.numberOfAlignedDiscs(opponent)

	if opponentFourAligned > 0 {
		return MinInt
	}
	return playerFourAligned*MaxInt + playerThreeAligned*100 + playerTwoAligned
}

func (board Board) numberOfAlignedDiscs(player uint8) (fourAligned, threeAligned, twoAligned int) {
	// convolution algorithm

	// vertical lines
	for i := 0; i < boardWidth; i++ {
		for j := 0; j < boardHeight-4; j++ {
			// part := board[j : j+4][i]
		}
	}
	return
}

func getOpponent(currentPlayer uint8) uint8 {
	if currentPlayer == 1 {
		return 2
	}
	return 1
}

func bestMove(board Board, player, depth int) (column, score int) {

	if depth == 0 {
		// compute column score
	}
	score = 0
	depth--
	// drop a disc on every column and compute possibility
	for i := 0; i <= boardWidth; i++ {
		// dropDisc(board, player, i)

		// columnScore := getColumnScore(board, player, depth)
		// if columnScore > score {
		// 	score = columnScore
		// 	column = i
		// }
	}
	return
}
