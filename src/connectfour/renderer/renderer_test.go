package renderer

import (
	"connectfour"
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
		`
    o
  x o x
  xxo o
 xxoxox
xoxxoxo`

	board := connectfour.New(boardInput)

	fmt.Println(Render(board))

	// Output:
	//
	//     o
	//   x o x
	//   xxo o
	//  xxoxox
	// xoxxoxo

}
