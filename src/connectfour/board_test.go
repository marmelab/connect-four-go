package connectfour

import (
	"testing"
)

func TestItPrintsTheRightBoardBack(t *testing.T) {
	boardInput := "0000000\n0000200\n0010201\n0011202\n0112121\n1211212"

	board := BoardFromString(boardInput)

	boardOutput := BoardToString(board)

	if boardInput != boardOutput {
		t.Error("Expected board output to be the same as input")
	}
}
