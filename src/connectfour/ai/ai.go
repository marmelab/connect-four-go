package ai

import (
	"connectfour/board"
	"errors"
	"math"
	"time"
)

const MinScore int = math.MinInt32

type ScoredBoard struct {
	CurrentBoard   board.Board
	CurrentScoring int
}

type BestMove struct {
	Column int
	Error  error
}

func HasWon(gameBoard board.Board, player int) bool {
	return numberOfAlignedDiscs(gameBoard, player, 4) > 0
}

func score(gameBoard board.Board, player int) int {
	fourSeriesNumber := numberOfAlignedDiscs(gameBoard, player, 4)
	threeSeriesNumber := numberOfAlignedDiscs(gameBoard, player, 3)
	twoSeriesNumber := numberOfAlignedDiscs(gameBoard, player, 2)

	opponent := getOpponent(player)
	opponentFourSeriesNumber := numberOfAlignedDiscs(gameBoard, opponent, 4)
	opponentThreeSeriesNumber := numberOfAlignedDiscs(gameBoard, opponent, 3)

	if opponentFourSeriesNumber > 0 {
		return -10000
	}

	return fourSeriesNumber*10000 + opponentThreeSeriesNumber*9000 + threeSeriesNumber*100 + twoSeriesNumber
}

func numberOfAlignedDiscs(gameBoard board.Board, player int, chunkSize int) int {
	count := 0

	for _, chunk := range getAllChunks(gameBoard, chunkSize) {
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

func getAllChunks(gameBoard board.Board, chunkSize int) [][]int {
	chunks := [][]int{}
	results := make(chan []int, 1)
	finished := make(chan bool, 4)

	go getHorizontalChunks(gameBoard, chunkSize, results, finished)
	go getVerticalChunks(gameBoard, chunkSize, results, finished)
	go getBottomLeftTopRightDiagonalChunks(gameBoard, chunkSize, results, finished)
	go getTopLeftBottomRightDiagonalChunks(gameBoard, chunkSize, results, finished)

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

func getHorizontalChunks(gameBoard board.Board, chunkSize int, results chan []int, finished chan bool) {
	for y := 0; y < board.BoardHeight; y++ {
		line := gameBoard[y]
		for x := 0; x < board.BoardWidth-chunkSize+1; x++ {
			results <- line[x : x+chunkSize]
		}
	}
	finished <- true
}

func getVerticalChunks(gameBoard board.Board, chunkSize int, results chan []int, finished chan bool) {
	for x := 0; x < board.BoardWidth; x++ {
		for y := 0; y < board.BoardHeight-chunkSize+1; y++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = gameBoard[y+z][x]
			}
			results <- part
		}
	}

	finished <- true
}

func getBottomLeftTopRightDiagonalChunks(gameBoard board.Board, chunkSize int, results chan []int, finished chan bool) {
	for x := 0; x < board.BoardWidth-chunkSize+1; x++ {
		for y := chunkSize - 1; y < board.BoardHeight; y++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = gameBoard[y-z][x+z]
			}
			results <- part
		}
	}

	finished <- true
}

func getTopLeftBottomRightDiagonalChunks(gameBoard board.Board, chunkSize int, results chan []int, finished chan bool) {
	for x := chunkSize - 1; x < board.BoardWidth; x++ {
		for y := chunkSize - 1; y < board.BoardHeight; y++ {
			part := make([]int, chunkSize)
			for z := 0; z < chunkSize; z++ {
				part[z] = gameBoard[y-z][x-z]
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

func NextBestMoveInTime(gameBoard board.Board, player int, duration time.Duration) (int, error) {
	results := make(chan BestMove, 1)
	timeout := make(chan bool, 1)
	column := -1
	var error error

	go func() {
		time.Sleep(duration)
		timeout <- true
	}()

	go NextBestMove(gameBoard, player, results)

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

func NextBestMove(gameBoard board.Board, player int, results chan BestMove) {
	currentPlayer := player
	defer close(results)
	var scoredBoards [][]ScoredBoard = make([][]ScoredBoard, board.BoardWidth)
	firstPossibleBoards := 0

	for i := 0; i < board.BoardWidth; i++ {
		scoredBoards[i] = make([]ScoredBoard, 0)

		nextBoard, err := gameBoard.AddDisc(i, currentPlayer)
		if err != nil {
			continue
		}

		if HasWon(nextBoard, player) {
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

		for i := 0; i < board.BoardWidth; i++ {
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
	var nextScoredBoardsByColumn [][]ScoredBoard = make([][]ScoredBoard, board.BoardWidth)

	for i := 0; i < board.BoardWidth; i++ {
		nextScoredBoardsByColumn[i] = make([]ScoredBoard, 0)

		nextScoredBoardsOneColumn := guessNextBoards(scoredBoards[i], currentPlayer, scoringPlayer)

		nextScoredBoardsByColumn[i] = append(nextScoredBoardsByColumn[i], nextScoredBoardsOneColumn...)
	}

	return nextScoredBoardsByColumn
}

func guessNextBoards(scoredBoards []ScoredBoard, currentPlayer, scoringPlayer int) []ScoredBoard {
	var nextScoredBoards []ScoredBoard

	for _, scoredBoard := range scoredBoards {
		for i := 0; i < board.BoardWidth; i++ {
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
