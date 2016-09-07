package connectfour

import (
	"strings"
)

const FirstPlayerChar string = "x"
const SecondPlayerChar string = "o"

type Board [6][7]uint8

func New(input string) Board {
	lines := strings.Split(input, "\n")
	board := Board{}
	for lineIndex, line := range lines {
		for charIndex, char := range line {
			switch string(char) {
			case FirstPlayerChar:
				board[lineIndex][charIndex] = 1
			case SecondPlayerChar:
				board[lineIndex][charIndex] = 2
			}
		}
	}
	return board
}
