package ai

import (
	"connectfour"
	"math"
)

const MinScore int = math.MinInt32

type ScoredBoard struct {
	currentBoard   connectfour.Board
	currentScoring int
}

func score(board connectfour.Board, player int) int {
	fourSeriesNumber := numberOfAlignedDiscs(board, player, 4)
	threeSeriesNumber := numberOfAlignedDiscs(board, player, 3)
	twoSeriesNumber := numberOfAlignedDiscs(board, player, 2)

	opponent := getOpponent(player)
	opponentFourSeriesNumber := numberOfAlignedDiscs(board, opponent, 4)

	if opponentFourSeriesNumber > 0 {
		return -10000
	}

	return fourSeriesNumber*10000 + threeSeriesNumber*100 + twoSeriesNumber
}

func numberOfAlignedDiscs(board connectfour.Board, player int, chunkSize int) int {
	count := 0

	for _, chunk := range getAllChunks(board, chunkSize) {
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

func getAllChunks(board connectfour.Board, chunkSize int) [][]int {
	chunks := [][]int{}

	horizontalChunks := getHorizontalChunks(board, chunkSize)
	chunks = append(chunks, horizontalChunks...)

	verticalChunks := getVerticalChunks(board, chunkSize)
	chunks = append(chunks, verticalChunks...)

	bottomLeftTopRightDiagonalChunks := getBottomLeftTopRightDiagonalChunks(board, chunkSize)
	chunks = append(chunks, bottomLeftTopRightDiagonalChunks...)

	topLeftBottomRightDiagonalChunks := getTopLeftBottomRightDiagonalChunks(board, chunkSize)
	chunks = append(chunks, topLeftBottomRightDiagonalChunks...)

	return chunks
}

func getHorizontalChunks(board connectfour.Board, chunkSize int) [][]int {
	chunks := [][]int{}

	for y := 0; y < connectfour.BoardHeight; y++ {
		line := board[y]
		for x := 0; x < connectfour.BoardWidth-chunkSize; x++ {
			chunks = append(chunks, line[x:x+chunkSize])
		}
	}

	return chunks
}

func getVerticalChunks(board connectfour.Board, chunkSize int) [][]int {
	chunks := [][]int{}

	for x := 0; x < connectfour.BoardWidth; x++ {
		for y := 0; y < connectfour.BoardHeight-chunkSize+1; y++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[y+z][x]
			}
			chunks = append(chunks, part)
		}
	}

	return chunks
}

func getBottomLeftTopRightDiagonalChunks(board connectfour.Board, chunkSize int) [][]int {
	chunks := [][]int{}

	for x := 0; x < connectfour.BoardWidth-chunkSize+1; x++ {
		for y := chunkSize - 1; y < connectfour.BoardHeight; y++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[y-z][x+z]
			}
			chunks = append(chunks, part)
		}
	}

	return chunks
}

func getTopLeftBottomRightDiagonalChunks(board connectfour.Board, chunkSize int) [][]int {
	chunks := [][]int{}

	for x := chunkSize - 1; x < connectfour.BoardWidth; x++ {
		for y := chunkSize - 1; y < connectfour.BoardHeight; y++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[y-z][x-z]
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

func NextBestMove(board connectfour.Board, player int) int {
	currentPlayer := player

	var scoredBoards [][]ScoredBoard = make([][]ScoredBoard, connectfour.BoardWidth)

	for i := 0; i < connectfour.BoardWidth; i++ {
		scoredBoards[i] = make([]ScoredBoard, 0)

		nextBoard, err := board.AddDisc(i, currentPlayer)
		if err != nil {
			continue
		}

		scoredBoard := ScoredBoard{
			currentBoard:   nextBoard,
			currentScoring: score(nextBoard, currentPlayer),
		}

		scoredBoards[i] = append(scoredBoards[i], scoredBoard)
	}
	depth := 1
	var bestColumn int
	for {
		bestScore := MinScore

		for i := 0; i < connectfour.BoardWidth; i++ {
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
	var nextScoredBoardsByColumn [][]ScoredBoard = make([][]ScoredBoard, connectfour.BoardWidth)

	for i := 0; i < connectfour.BoardWidth; i++ {
		nextScoredBoardsByColumn[i] = make([]ScoredBoard, 0)

		nextScoredBoardsOneColumn := guessNextBoards(scoredBoards[i], currentPlayer, scoringPlayer)

		nextScoredBoardsByColumn[i] = append(nextScoredBoardsByColumn[i], nextScoredBoardsOneColumn...)
	}

	return nextScoredBoardsByColumn
}

func guessNextBoards(scoredBoards []ScoredBoard, currentPlayer, scoringPlayer int) []ScoredBoard {
	var nextScoredBoards []ScoredBoard

	for _, scoredBoard := range scoredBoards {
		for i := 0; i < connectfour.BoardWidth; i++ {
			nextBoard, err := scoredBoard.currentBoard.AddDisc(i, currentPlayer)
			if err != nil {
				continue
			}

			nextScoredBoard := ScoredBoard{
				currentBoard:   nextBoard,
				currentScoring: score(nextBoard, scoringPlayer),
			}

			nextScoredBoards = append(nextScoredBoards, nextScoredBoard)
		}
	}

	return nextScoredBoards
}
