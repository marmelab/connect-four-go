package board

import (
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
