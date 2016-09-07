package connectfour

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const boardWidth int = 7
const boardHeight int = 6

const FirstPlayerChar string = "x"
const SecondPlayerChar string = "o"

const MinScore int = math.MinInt32

type Board [boardHeight][boardWidth]int

type ScoredBoard struct {
	currentBoard   Board // currentBoard
	currentScoring int   //currentScoring
}

func New(input string) Board {
	lines := strings.Split(input, "\n")
	board := Board{}

	for lineIndex, line := range lines {
		for i := 0; i < boardWidth; i++ {
			var char string
			if i < len(line) {
				char = string(line[i])
			} else {
				char = ""
			}
			switch char {
			case FirstPlayerChar:
				board[lineIndex][i] = 1
			case SecondPlayerChar:
				board[lineIndex][i] = 2
			default:
				board[lineIndex][i] = 0
			}
		}
	}

	return board
}

func (board Board) AddDisc(column, player int) (Board, error) {
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

func (board Board) score(player int) int {

	playerFourAligned := board.numberOfAlignedDiscs(player, 4)
	playerThreeAligned := board.numberOfAlignedDiscs(player, 3)
	playerTwoAligned := board.numberOfAlignedDiscs(player, 2)

	opponent := getOpponent(player)
	opponentFourAligned := board.numberOfAlignedDiscs(opponent, 4)

	if opponentFourAligned > 0 {
		return -10000
	}
	return playerFourAligned*10000 + playerThreeAligned*100 + playerTwoAligned
}

func (board Board) numberOfAlignedDiscs(player int, chunkSize int) int {
	count := 0

	for _, chunk := range board.getAllChunks(chunkSize) {
		if areConsecutives(chunk, player) {
			count++
		}
	}

	return count
}

func areConsecutives(cells []int, player int) bool {
	for i := 0; i < len(cells)-1; i++ {
		if cells[i] != cells[i+1] || cells[i] != player {
			return false
		}
	}
	return true
}

func (board Board) getAllChunks(chunkSize int) [][]int {
	chunks := [][]int{}

	// horizontal
	for x := 0; x < boardHeight; x++ {
		line := board[x]
		for y := 0; y < boardWidth-chunkSize; y++ {
			chunks = append(chunks, line[y:y+chunkSize])
		}
	}

	//vertical
	for y := 0; y < boardWidth; y++ {
		for x := 0; x < boardHeight-chunkSize+1; x++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[x+z][y]
			}
			chunks = append(chunks, part)
		}
	}

	// diagonals /
	for y := 0; y < boardWidth-chunkSize+1; y++ {
		for x := chunkSize - 1; x < boardHeight; x++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[x-z][y+z]
			}
			chunks = append(chunks, part)
		}
	}

	// diagonals \
	for y := chunkSize - 1; y < boardWidth; y++ {
		for x := chunkSize - 1; x < boardHeight; x++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[x-z][y-z]
			}
			chunks = append(chunks, part)
		}
	}

	return chunks
}

func getOpponent(currentPlayer int) int {
	if currentPlayer == 1 {
		return 2
	}
	return 1
}

func (board Board) NextBestMove(player int) int {
	currentPlayer := player

	var scoredBoards [][]ScoredBoard = make([][]ScoredBoard, boardWidth)

	for i := 0; i < boardWidth; i++ {
		scoredBoards[i] = make([]ScoredBoard, 0)

		nextBoard, err := board.AddDisc(i, currentPlayer)
		if err != nil {
			continue
		}

		scoredBoard := ScoredBoard{
			currentBoard:   nextBoard,
			currentScoring: nextBoard.score(currentPlayer),
		}

		scoredBoards[i] = append(scoredBoards[i], scoredBoard)
	}
	depth := 1
	var bestColumn int
	for depth < 4 {
		bestScore := MinScore

		for i := 0; i < boardWidth; i++ {
			score := aggregateScoring(scoredBoards[i])

			if score > bestScore {
				bestScore = score
				bestColumn = i
			}
		}

		fmt.Println("Depth", depth, "Best Score", bestScore, "Best column", bestColumn)

		currentPlayer = getOpponent(currentPlayer)

		scoredBoards = guessNextBoardsAggregated(scoredBoards, currentPlayer, player)
	}

	return bestColumn
	//
	// bestScore := MinScore
	// var bestColumn int
	// for i := 0; i <= boardWidth; i++ {
	// 	score, err := board.scan(i, player, 3)
	// 	if err != nil {
	// 		continue
	// 	}
	// 	if score > bestScore {
	// 		bestScore = score
	// 		bestColumn = i
	// 	}
	// }
	// return bestColumn
}

func aggregateScoring(scoredBoards []ScoredBoard) int {
	score := 0
	for _, scoredBoard := range scoredBoards {
		score += scoredBoard.currentScoring
	}
	return score
}

func guessNextBoardsAggregated(scoredBoards [][]ScoredBoard, currentPlayer, scoringPlayer int) [][]ScoredBoard {
	var nextScoredBoardsByColumn [][]ScoredBoard = make([][]ScoredBoard, boardWidth)

	for i := 0; i < boardWidth; i++ {
		nextScoredBoardsByColumn[i] = make([]ScoredBoard, 0)

		nextScoredBoardsOneColumn := guessNextBoards(scoredBoards[i], currentPlayer, scoringPlayer)

		nextScoredBoardsByColumn[i] = append(nextScoredBoardsByColumn[i], nextScoredBoardsOneColumn...)
	}

	return nextScoredBoardsByColumn
}

func guessNextBoards(scoredBoards []ScoredBoard, currentPlayer, scoringPlayer int) []ScoredBoard {
	var nextScoredBoards []ScoredBoard

	for _, scoredBoard := range scoredBoards {
		for i := 0; i < boardWidth; i++ {
			nextBoard, err := scoredBoard.currentBoard.AddDisc(i, currentPlayer)
			if err != nil {
				continue
			}

			nextScoredBoard := ScoredBoard{
				currentBoard:   nextBoard,
				currentScoring: nextBoard.score(scoringPlayer),
			}

			nextScoredBoards = append(nextScoredBoards, nextScoredBoard)
		}
	}

	return nextScoredBoards
}
