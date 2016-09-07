package connectfour

import (
	"fmt"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestPlayerCanDropADisc(t *testing.T) {
	board := Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 1, 0, 2, 0, 1},
		{0, 0, 1, 1, 2, 0, 2},
		{0, 1, 1, 2, 1, 2, 1},
		{1, 2, 1, 1, 2, 1, 2},
	}

	board, err := board.AddDisc(0, 2)
	check(err)

	if board[4][0] != 2 {
		t.Error("Expected 2, got ", board[4][0])
	}
}

func TestNumberOfAlignedDiscs(t *testing.T) {
	board := Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 2, 0, 0},
		{0, 0, 1, 0, 2, 0, 1},
		{0, 0, 1, 1, 2, 0, 2},
		{0, 1, 1, 2, 1, 2, 1},
		{1, 2, 1, 1, 2, 1, 2},
	}

	numberOfAlignedDiscs := board.numberOfAlignedDiscs(1, 3)

	if numberOfAlignedDiscs != 5 {
		t.Error("Expected 5, got ", numberOfAlignedDiscs)
	}
}

func TestScoreSecondPlayerShouldBeHigher(t *testing.T) {
	board := Board{
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	firstPlayerScore := board.score(1)
	secondPlayerScore := board.score(2)

	if firstPlayerScore > secondPlayerScore {
		t.Error("Expected second player score to be higher than first player score")
	}
}

func TestGuessNextBoards(t *testing.T) {
	board := Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	scoredBoards := make([]ScoredBoard, 0)
	scoredBoards = append(scoredBoards, ScoredBoard{
		currentBoard: board,
	})

	nextScoredBoards := guessNextBoards(scoredBoards, 1, 1)
	fmt.Println(nextScoredBoards)
	firstBoard := Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{1, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	if nextScoredBoards[0].currentBoard != firstBoard {
		t.Error("Expected boards to contain next board playing with first column")
	}

	secondBoard := Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 1, 2, 0, 0, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	if nextScoredBoards[1].currentBoard != secondBoard {
		t.Error("Expected boards to contain next board playing with second column")
	}

	if len(nextScoredBoards) > 6 {
		t.Error("Expected next boards not to contain any board on column two")
	}

	thirdBoard := Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 1, 0, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	if nextScoredBoards[2].currentBoard != thirdBoard {
		t.Error("Expected boards to contain next board playing with fourth column")
	}

	fourthBoard := Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 1, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	if nextScoredBoards[3].currentBoard != fourthBoard {
		t.Error("Expected boards to contain next board playing with fifth column")
	}

	fifthBoard := Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 1, 1, 2, 1, 1, 0},
		{1, 1, 1, 2, 2, 2, 0},
	}

	if nextScoredBoards[4].currentBoard != fifthBoard {
		t.Error("Expected boards to contain next board playing with sixth column")
	}

	sixthBoard := Board{
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0},
		{0, 0, 2, 0, 0, 0, 0},
		{0, 1, 1, 2, 1, 0, 0},
		{1, 1, 1, 2, 2, 2, 1},
	}

	if nextScoredBoards[5].currentBoard != sixthBoard {
		t.Error("Expected boards to contain next board playing with seventh column")
	}
}
