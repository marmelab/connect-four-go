package parser

import (
	"connectfour"
	"strings"
)

func Parse(input string, firstPlayerChar, secondPlayerChar, emptyCellChar string) connectfour.Board {
	lines := strings.Split(input, "\n")
	board := connectfour.Board{}

	for lineIndex, line := range lines {
		for i := 0; i < connectfour.BoardWidth; i++ {
			var char string
			if i < len(line) {
				char = string(line[i])
			} else {
				char = emptyCellChar
			}
			switch char {
			case firstPlayerChar:
				board[lineIndex][i] = 1
			case secondPlayerChar:
				board[lineIndex][i] = 2
			default:
				board[lineIndex][i] = 0
			}
		}
	}

	return board
}
