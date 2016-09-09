package renderer

import (
	"connectfour/parser"
)

func ExampleItPrintsTheRightBoardBack() {
	boardInput :=
		`.......
....o..
..x.o.x
..xxo.o
.xxoxox
xoxxoxo`

	board := parser.Parse(boardInput, "x", "o", ".")

	Render(board, "x", "o", ".", "%v")

	// Output:
	// 1234567
	// .......
	// ....o..
	// ..x.o.x
	// ..xxo.o
	// .xxoxox
	// xoxxoxo
}
