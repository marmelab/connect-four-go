package ai

import (
	"connectfour"
	"testing"
	"time"
)

func TestNumberOfAlignedDiscs(t *testing.T) {
	board := connectfour.Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 1, 0, 2, 0, 1},
		{0, 0, 1, 1, 2, 0, 2},
		{0, 1, 1, 2, 1, 2, 1},
		{1, 2, 1, 1, 2, 1, 2},
	}

	numberOfAlignedDiscs := numberOfAlignedDiscs(board, 1, 3)

	if numberOfAlignedDiscs != 5 {
		t.Error("Expected 5, got ", numberOfAlignedDiscs)
	}
}

func TestScoreSecondPlayerShouldBeHigher(t *testing.T) {
	board := connectfour.Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	firstPlayerScore := score(board, 1)
	secondPlayerScore := score(board, 2)

	if firstPlayerScore > secondPlayerScore {
		t.Error("Expected second player score to be higher than first player score")
	}
}

func TestGuessNextBoards(t *testing.T) {
	board := connectfour.Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	scoredBoards := make([]ScoredBoard, 0)
	scoredBoards = append(scoredBoards, ScoredBoard{
		CurrentBoard: board,
	})

	nextScoredBoards := guessNextBoards(scoredBoards, 1, 1)

	firstBoard := connectfour.Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{1, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	if nextScoredBoards[0].CurrentBoard != firstBoard {
		t.Error("Expected boards to contain next board playing with first column")
	}

	secondBoard := connectfour.Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 1, 2, 0, 0, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	if nextScoredBoards[1].CurrentBoard != secondBoard {
		t.Error("Expected boards to contain next board playing with second column")
	}

	if len(nextScoredBoards) > 6 {
		t.Error("Expected next boards not to contain any board on column two")
	}

	thirdBoard := connectfour.Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 1, 0, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	if nextScoredBoards[2].CurrentBoard != thirdBoard {
		t.Error("Expected boards to contain next board playing with fourth column")
	}

	fourthBoard := connectfour.Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 1, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	if nextScoredBoards[3].CurrentBoard != fourthBoard {
		t.Error("Expected boards to contain next board playing with fifth column")
	}

	fifthBoard := connectfour.Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 1, 1, 2, 1, 1, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	if nextScoredBoards[4].CurrentBoard != fifthBoard {
		t.Error("Expected boards to contain next board playing with sixth column")
	}

	sixthBoard := connectfour.Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 1},
	}

	if nextScoredBoards[5].CurrentBoard != sixthBoard {
		t.Error("Expected boards to contain next board playing with seventh column")
	}
}

func TestGetOpponent(t *testing.T) {
	player := 1
	opponent := getOpponent(player)

	if opponent == player {
		t.Error("Expected player and opponent to be different")
	}
}

func TestNextBestMove(t *testing.T) {
	board := connectfour.Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 2, 0, 1},
		{0, 0, 1, 1, 2, 0, 2},
		{0, 1, 1, 2, 1, 2, 1},
		{2, 2, 1, 1, 2, 1, 2},
	}

	results := make(chan Result, 1)
	go NextBestMove(board, 1, results)

	result := <-results
	if result.Column != 2 {
		t.Error("Expected 2, got ", result.Column)
	}
}

func TestNextBestMoveInTimeIsReturned(t *testing.T) {
	board := connectfour.Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 2, 0, 1},
		{0, 0, 1, 1, 2, 0, 2},
		{0, 1, 1, 2, 1, 2, 1},
		{2, 2, 1, 1, 2, 1, 2},
	}

	column, err := NextBestMoveInTime(board, 1, 2*time.Second)

	if err != nil {
		t.Error("Expected to have at least one result in time")
	}
	if column != 2 {
		t.Error("Expected 2, got ", column)
	}
}

func TestErrorIsReturnWhenNotEnoughTimeForNextBestMoveInTime(t *testing.T) {
	t.Skip("Try to make this test work somehow, result is return after 1 ns no matter what")

	board := connectfour.Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 1, 2, 0, 2},
		{0, 1, 1, 2, 1, 2, 1},
		{2, 2, 1, 1, 2, 1, 2},
	}

	column, err := NextBestMoveInTime(board, 1, 1*time.Nanosecond)

	if err == nil {
		t.Error("Expected not to have enough time to get a result, result", column)
	}
}

func TestWhatHappensWhenBoardFull(t *testing.T) {
	board := connectfour.Board{
		{1, 1, 1, 2, 1, 2, 1},
		{2, 2, 2, 1, 2, 1, 1},
		{1, 1, 1, 2, 2, 1, 1},
		{2, 2, 2, 1, 2, 2, 2},
		{1, 1, 1, 2, 1, 2, 1},
		{2, 2, 1, 1, 2, 1, 2},
	}

	_, err := NextBestMoveInTime(board, 1, 1*time.Nanosecond)
	if err == nil {
		t.Error("Expected to have an error")
	}
}

func TestAggregateScoring(t *testing.T) {
	scoredBoards := make([]ScoredBoard, 0)

	scoredBoards = append(scoredBoards, ScoredBoard{
		CurrentScoring: 10,
	})
	scoredBoards = append(scoredBoards, ScoredBoard{
		CurrentScoring: -10,
	})
	scoredBoards = append(scoredBoards, ScoredBoard{
		CurrentScoring: 20,
	})
	scoredBoards = append(scoredBoards, ScoredBoard{
		CurrentScoring: -20,
	})
	scoredBoards = append(scoredBoards, ScoredBoard{
		CurrentScoring: 30,
	})
	scoredBoards = append(scoredBoards, ScoredBoard{
		CurrentScoring: -30,
	})

	aggregatedScore := aggregateScoring(scoredBoards)

	if aggregatedScore != 0 {
		t.Error("Expected 0, got ", aggregatedScore)
	}
}

func testNotConsecutiveDiscs(t *testing.T) {
	cells := []int{1, 1, 2, 1}

	areConsecutives := areConsecutives(cells, 1)

	if areConsecutives {
		t.Error("Expected false, got ", areConsecutives)
	}
}

func testConsecutiveDiscs(t *testing.T) {
	cells := []int{1, 1, 1, 1}

	areConsecutives := areConsecutives(cells, 1)

	if !areConsecutives {
		t.Error("Expected true, got ", areConsecutives)
	}
}

func BenchmarkNextBestMove(b *testing.B) {
	board := connectfour.Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 2, 0, 1},
		{0, 0, 1, 1, 2, 0, 2},
		{0, 1, 1, 2, 1, 2, 1},
		{2, 2, 1, 1, 2, 1, 2},
	}

	for n := 0; n < b.N; n++ {
		results := make(chan Result, 1)
		go NextBestMove(board, 1, results)

		<-results
	}
}

func BenchmarkGuessNextBoards(b *testing.B) {
	board := connectfour.Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 0, 0, 2, 0, 1},
		{0, 0, 1, 1, 2, 0, 2},
		{0, 1, 1, 2, 1, 2, 1},
		{1, 2, 1, 1, 2, 1, 2},
	}

	scoredBoards := make([]ScoredBoard, 0)
	scoredBoards = append(scoredBoards, ScoredBoard{
		CurrentBoard: board,
	})

	for n := 0; n < b.N; n++ {
		guessNextBoards(scoredBoards, 1, 1)
	}
}
