package connectfour

import (
	"errors"
	"math"
)

const BoardWidth int = 7
const BoardHeight int = 6

const MinScore int = math.MinInt32

type Board [BoardHeight][BoardWidth]int

type ScoredBoard struct {
	currentBoard   Board
	currentScoring int
}

func (board Board) AddDisc(column, player int) (Board, error) {
	dropped := false

	for i := BoardHeight - 1; i >= 0; i-- {
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
	fourSeriesNumber := board.numberOfAlignedDiscs(player, 4)
	threeSeriesNumber := board.numberOfAlignedDiscs(player, 3)
	twoSeriesNumber := board.numberOfAlignedDiscs(player, 2)

	opponent := getOpponent(player)
	opponentFourSeriesNumber := board.numberOfAlignedDiscs(opponent, 4)

	if opponentFourSeriesNumber > 0 {
		return -10000
	}

	return fourSeriesNumber*10000 + threeSeriesNumber*100 + twoSeriesNumber
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
	for x := 0; x < BoardHeight; x++ {
		line := board[x]
		for y := 0; y < BoardWidth-chunkSize; y++ {
			chunks = append(chunks, line[y:y+chunkSize])
		}
	}

	//vertical
	for y := 0; y < BoardWidth; y++ {
		for x := 0; x < BoardHeight-chunkSize+1; x++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[x+z][y]
			}
			chunks = append(chunks, part)
		}
	}

	// diagonals /
	for y := 0; y < BoardWidth-chunkSize+1; y++ {
		for x := chunkSize - 1; x < BoardHeight; x++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[x-z][y+z]
			}
			chunks = append(chunks, part)
		}
	}

	// diagonals \
	for y := chunkSize - 1; y < BoardWidth; y++ {
		for x := chunkSize - 1; x < BoardHeight; x++ {
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

	var scoredBoards [][]ScoredBoard = make([][]ScoredBoard, BoardWidth)

	for i := 0; i < BoardWidth; i++ {
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
	for {
		bestScore := MinScore

		for i := 0; i < BoardWidth; i++ {
			score := aggregateScoring(scoredBoards[i])

			if score > bestScore {
				bestScore = score
				bestColumn = i
			}
		}

		if depth == 5 {
			break
		}

		currentPlayer = getOpponent(currentPlayer)

		scoredBoards = guessNextBoardsAggregated(scoredBoards, currentPlayer, player)

		depth++
	}

	return bestColumn
}

func aggregateScoring(scoredBoards []ScoredBoard) int {
	score := 0
	for _, scoredBoard := range scoredBoards {
		score += scoredBoard.currentScoring
	}
	return score
}

func guessNextBoardsAggregated(scoredBoards [][]ScoredBoard, currentPlayer, scoringPlayer int) [][]ScoredBoard {
	var nextScoredBoardsByColumn [][]ScoredBoard = make([][]ScoredBoard, BoardWidth)

	for i := 0; i < BoardWidth; i++ {
		nextScoredBoardsByColumn[i] = make([]ScoredBoard, 0)

		nextScoredBoardsOneColumn := guessNextBoards(scoredBoards[i], currentPlayer, scoringPlayer)

		nextScoredBoardsByColumn[i] = append(nextScoredBoardsByColumn[i], nextScoredBoardsOneColumn...)
	}

	return nextScoredBoardsByColumn
}

func guessNextBoards(scoredBoards []ScoredBoard, currentPlayer, scoringPlayer int) []ScoredBoard {
	var nextScoredBoards []ScoredBoard

	for _, scoredBoard := range scoredBoards {
		for i := 0; i < BoardWidth; i++ {
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
