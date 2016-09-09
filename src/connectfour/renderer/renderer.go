package renderer

import (
	"bytes"
	"connectfour/board"
	"fmt"
	"strconv"
)

func Render(gameBoard board.Board, firstPlayerChar, secondPlayerChar, emptyCellChar, headerChar string) string {
	var buffer bytes.Buffer
	for x := 1; x <= board.BoardWidth; x++ {
		buffer.WriteString(fmt.Sprintf(headerChar, strconv.Itoa(x)))
	}
	buffer.WriteString("\n")
	for y := 0; y < board.BoardHeight; y++ {
		for x := 0; x < board.BoardWidth; x++ {
			switch gameBoard[y][x] {
			case 1:
				buffer.WriteString(firstPlayerChar)
			case 2:
				buffer.WriteString(secondPlayerChar)
			default:
				buffer.WriteString(emptyCellChar)
			}
		}
		buffer.WriteString("\n")
	}
	return buffer.String()
}
