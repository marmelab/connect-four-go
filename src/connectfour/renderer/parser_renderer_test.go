package renderer

import (
	"connectfour/parser"
	"fmt"
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

	fmt.Println(Render(board, "x", "o", "."))

	// Output:
	// .......
	// ....o..
	// ..x.o.x
	// ..xxo.o
	// .xxoxox
	// xoxxoxo
}
