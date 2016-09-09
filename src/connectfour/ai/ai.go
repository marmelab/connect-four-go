package ai

import (
	"connectfour"
	"errors"
	"math"
	"time"
)

const MinScore int = math.MinInt32

type ScoredBoard struct {
	CurrentBoard   connectfour.Board
	CurrentScoring int
}

type BestMove struct {
	Column int
	Error  error
}

func hasWon(board connectfour.Board, player int) bool {
	return numberOfAlignedDiscs(board, player, 4) > 0
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
	results := make(chan []int, 1)
	finished := make(chan bool, 4)

	go getHorizontalChunks(board, chunkSize, results, finished)
	go getVerticalChunks(board, chunkSize, results, finished)
	go getBottomLeftTopRightDiagonalChunks(board, chunkSize, results, finished)
	go getTopLeftBottomRightDiagonalChunks(board, chunkSize, results, finished)

	nbFinished := 0
	for nbFinished < 4 {
		select {
		case <-finished:
			nbFinished++
			break
		case parts := <-results:
			chunks = append(chunks, parts)
		}
	}
	return chunks
}

func getHorizontalChunks(board connectfour.Board, chunkSize int, results chan []int, finished chan bool) {
	for y := 0; y < connectfour.BoardHeight; y++ {
		line := board[y]
		for x := 0; x < connectfour.BoardWidth-chunkSize+1; x++ {
			results <- line[x : x+chunkSize]
		}
	}
	finished <- true
}

func getVerticalChunks(board connectfour.Board, chunkSize int, results chan []int, finished chan bool) {
	for x := 0; x < connectfour.BoardWidth; x++ {
		for y := 0; y < connectfour.BoardHeight-chunkSize+1; y++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[y+z][x]
			}
			results <- part
		}
	}

	finished <- true
}

func getBottomLeftTopRightDiagonalChunks(board connectfour.Board, chunkSize int, results chan []int, finished chan bool) {
	for x := 0; x < connectfour.BoardWidth-chunkSize+1; x++ {
		for y := chunkSize - 1; y < connectfour.BoardHeight; y++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[y-z][x+z]
			}
			results <- part
		}
	}

	finished <- true
}

func getTopLeftBottomRightDiagonalChunks(board connectfour.Board, chunkSize int, results chan []int, finished chan bool) {
	for x := chunkSize - 1; x < connectfour.BoardWidth; x++ {
		for y := chunkSize - 1; y < connectfour.BoardHeight; y++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = board[y-z][x-z]
			}
			results <- part
		}
	}

	finished <- true
}

func getOpponent(currentPlayer int) int {
	if currentPlayer == 1 {
		return 2
	}
	return 1
}

func NextBestMoveInTime(board connectfour.Board, player int, duration time.Duration) (int, error) {
	results := make(chan BestMove, 1)
	timeout := make(chan bool, 1)
	column := -1
	var error error

	go func() {
		time.Sleep(duration)
		timeout <- true
	}()

	go NextBestMove(board, player, results)

	finished := false

	for !finished {
		select {
		case finished = <-timeout:
		case result, channelStillOpen := <-results:
			if !channelStillOpen {
				finished = true
				break
			}
			column = result.Column
			if result.Error != nil {
				error = result.Error
				break
			}
		}
	}

	if error != nil {
		return column, error
	}

	if column == -1 {
		return column, errors.New("No result found in time")
	}

	return column, nil
}

func NextBestMove(board connectfour.Board, player int, results chan BestMove) {
	currentPlayer := player
	defer close(results)
	var scoredBoards [][]ScoredBoard = make([][]ScoredBoard, connectfour.BoardWidth)
	firstPossibleBoards := 0

	for i := 0; i < connectfour.BoardWidth; i++ {
		scoredBoards[i] = make([]ScoredBoard, 0)

		nextBoard, err := board.AddDisc(i, currentPlayer)
		if err != nil {
			continue
		}

		if hasWon(nextBoard, player) {
			results <- BestMove{
				Column: i,
				Error:  nil,
			}
			return
		}

		scoredBoard := ScoredBoard{
			CurrentBoard:   nextBoard,
			CurrentScoring: score(nextBoard, currentPlayer),
		}

		scoredBoards[i] = append(scoredBoards[i], scoredBoard)
		firstPossibleBoards++
	}

	if firstPossibleBoards == 0 {
		results <- BestMove{
			Column: -1,
			Error:  errors.New("No possible move"),
		}
		return
	}

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
		results <- BestMove{
			Column: bestColumn,
			Error:  nil,
		}

		currentPlayer = getOpponent(currentPlayer)

		scoredBoards = guessNextBoardsAggregated(scoredBoards, currentPlayer, player)
	}
}

func aggregateScoring(scoredBoards []ScoredBoard) int {
	score := 0
	for _, scoredBoard := range scoredBoards {
		score += scoredBoard.CurrentScoring
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
			nextBoard, err := scoredBoard.CurrentBoard.AddDisc(i, currentPlayer)
			if err != nil {
				continue
			}

			nextScoredBoard := ScoredBoard{
				CurrentBoard:   nextBoard,
				CurrentScoring: score(nextBoard, scoringPlayer),
			}

			nextScoredBoards = append(nextScoredBoards, nextScoredBoard)
		}
	}

	return nextScoredBoards
}
