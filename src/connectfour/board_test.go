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

func TestItPrintsTheRightBoardBack(t *testing.T) {
	boardInput :=
		`0000000
0000200
0010201
0011202
0112121
1211212`

	board := BoardFromString(boardInput)

	fmt.Println(BoardToString(board))

	// Output:
	// 0000000
	// 0000200
	// 0010201
	// 0011202
	// 0112121
	// 1211212
}
}
