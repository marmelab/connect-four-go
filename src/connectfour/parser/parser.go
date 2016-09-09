package parser

import (
	"connectfour/board"
	"strings"
)

func Parse(input string, firstPlayerChar, secondPlayerChar, emptyCellChar string) board.Board {
	lines := strings.Split(input, "\n")
	gameBoard := board.Board{}
	for lineIndex, line := range lines {
		if lineIndex >= board.BoardHeight {
			break
		}
		for i := 0; i < board.BoardWidth; i++ {
			var char string
			if i < len(line) {
				char = string(line[i])
			} else {
				char = emptyCellChar
			}
			switch char {
			case firstPlayerChar:
				gameBoard[lineIndex][i] = 1
			case secondPlayerChar:
				gameBoard[lineIndex][i] = 2
			default:
				gameBoard[lineIndex][i] = 0
			}
		}
	}

	return gameBoard
}
