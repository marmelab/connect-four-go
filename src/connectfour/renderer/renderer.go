package renderer

import (
	"connectfour/board"
	"fmt"
	"strconv"
)

func Render(gameBoard board.Board, firstPlayerChar, secondPlayerChar, emptyCellChar, headerChar string) {
	for x := 1; x <= board.BoardWidth; x++ {
		fmt.Printf(headerChar, strconv.Itoa(x))
	}
	fmt.Println()
	for y := 0; y < board.BoardHeight; y++ {
		for x := 0; x < board.BoardWidth; x++ {
			switch gameBoard[y][x] {
			case 1:
				fmt.Print(firstPlayerChar)
			case 2:
				fmt.Print(secondPlayerChar)
			default:
				fmt.Print(emptyCellChar)
			}
		}
		fmt.Println()
	}
}
