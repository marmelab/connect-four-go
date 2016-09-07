package test

import (
	"connectfour"
	"connectfour/renderer"
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

	fmt.Println(renderer.ToString(board))

	// Output:
	//
	//     o
	//   x o x
	//   xxo o
	//  xxoxox
	// xoxxoxo

}

func TestPlayerCanDropADisc(t *testing.T) {
	boardInput :=
		`
    o
  x o x
  xxo o
 xxoxox
xoxxoxo`

	board := connectfour.New(boardInput)

	board, err := board.AddDisc(0, 2)
	check(err)

	fmt.Println(renderer.ToString(board))

	// Output:
	//
	//     o
	//   x o x
	//   xxo o
	// oxxoxox
	// xoxxoxo
}
